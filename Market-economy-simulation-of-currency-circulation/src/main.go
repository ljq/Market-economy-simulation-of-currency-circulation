package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 银行模拟
type Bank struct {
	reserveRatio  float64 // 存款准备金率
	totalDeposit  float64 // 存款总额
	availableLoan float64 // 可用信贷资金
	totalLoan     float64 // 贷款总额
	interestRate  float64 // 贷款利率
	repaymentRate float64 // 还款比例
}

// 市场模拟
type Market struct {
	bank       *Bank       // 银行
	people     []*Person   // 人员列表
	industries []*Industry // 行业列表
	totalMoney float64     // 货币总量
	m2Money    float64     // M2货币量
}

// 人员模拟
type Person struct {
	money    float64 // 拥有的货币数量
	wealth   float64 // 拥有的总财富，包括货币和商品
	loan     float64 // 债务金额
	repaying bool    // 是否正在还款
}

// 行业模拟
type Industry struct {
	products   float64 // 商品数量
	price      float64 // 商品价格
	profitRate float64 // 利润率
}

func main() {
	// 初始化随机数生成器
	rand.Seed(0)

	// 初始化银行
	bank := &Bank{
		reserveRatio:  0.1,
		interestRate:  0.05,
		repaymentRate: 0.1,
	}

	// 初始化市场
	market := &Market{
		bank:       bank,
		people:     make([]*Person, 1000),
		industries: make([]*Industry, 20),
		totalMoney: 10000,
		m2Money:    0,
	}

	// 初始化人员
	for i := 0; i < 1000; i++ {
		market.people[i] = &Person{
			money:    market.totalMoney / 1000,
			wealth:   0,
			loan:     0,
			repaying: false,
		}
	}

	// 初始化行业
	for i := 0; i < 20; i++ {
		market.industries[i] = &Industry{
			products:   1000,
			price:      1,
			profitRate: 0.1,
		}
	}

	// 开始模拟经济活动
	for {
		// 银行向市场提供信贷资金
		if market.bank.availableLoan == 0 && market.bank.totalDeposit > 0 {
			market.bank.availableLoan = market.bank.totalDeposit * (1 - market.bank.reserveRatio)
			market.bank.totalLoan += market.bank.availableLoan
			market.bank.totalDeposit -= market.bank.availableLoan
		}

		// 模拟商品交易
		for i := 0; i < 10000; i++ {
			// 随机选择两个不同的行业
			industry1 := market.industries[rand.Intn(20)]
			industry2 := market.industries[rand.Intn(20)]
			for industry1 == industry2 {
				industry2 = market.industries[rand.Intn(20)]
			}

			// 随机选择一个销售者和一个购买者
			seller := market.people[rand.Intn(1000)]
			buyer := market.people[rand.Intn(1000)]

			// 计算交易价格和利润
			price := (industry1.price + industry2.price) / 2
			profit := (price - industry1.price) * industry1.products * industry1.profitRate

			// 如果购买者拥有足够的货币，则进行交易
			if buyer.money >= price {
				seller.money += price - profit
				buyer.money -= price
				seller.wealth += price - profit
				buyer.wealth += profit
				industry1.products--
				industry2.products++
			}
		}

		// 模拟贷款发放和还款
		for _, person := range market.people {
			// 如果该人员属于市场中头部10%的人，并且没有债务，则可以申请贷款
			if person.wealth >= getTopWealthThreshold(market.people, 0.1) && person.loan == 0 {
				loanAmount := market.bank.availableLoan / 10
				if loanAmount > 0 {
					person.loan += loanAmount
					person.money += loanAmount
					market.bank.totalLoan += loanAmount
					market.bank.availableLoan -= loanAmount
				}
			}

			// 如果该人员拥有债务，则需要进行还款
			if person.loan > 0 {
				if !person.repaying {
					person.money -= person.loan * market.bank.repaymentRate
					person.repaying = true
				} else {
					interest := person.loan * market.bank.interestRate
					person.money -= interest
					person.loan -= person.loan * market.bank.repaymentRate
					if person.loan <= 0 {
						market.bank.totalDeposit += person.loan
						market.bank.totalLoan -= person.loan
						person.loan = 0
						person.repaying = false
					}
				}
			}
		}

		// 判断是否结束模拟
		if market.bank.availableLoan == 0 && market.bank.totalLoan == 0 {
			break
		}
	}

	// 计算最终的M2货币量和每个人拥有的实际货币总量
	for _, person := range market.people {
		market.m2Money += person.money + person.loan
	}

	// 输出结果
	fmt.Printf("最终的M2货币量：%v\n", market.m2Money)
	for i, person := range market.people {
		fmt.Printf("第 %d 个人的实际货币总量：%v\n", i+1, person.money+person.loan)
	}
}

// 获取前百分比的人所需的财富门槛
func getTopWealthThreshold(people []*Person, percentage float64) float64 {
	thresholdIndex := int(float64(len(people)) * percentage)
	sortedPeople := make([]*Person, len(people))
	copy(sortedPeople, people)
	sort.Slice(sortedPeople, func(i, j int) bool {
		return sortedPeople[i].wealth > sortedPeople[j].wealth
	})
	return sortedPeople[thresholdIndex].wealth
}
