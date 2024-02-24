### Algorithm simple preset basic conditions


Golang implements a monetary credit and market economy activity scenario, with basic parameters: the total issuance amount of the base currency M1 is 10000 yuan, one bank simulates a bank reserve requirement ratio of 10%, ignoring tax and fee issues. A simulated market has 1000 individuals and 20 industries with zero wealth, with personnel evenly distributed. The top 10% of people in each industry can apply for loans, and credit funds can be used until they are fully released to non lending. The market currency is fully circulated, Equal profit margins across industries; The initial cost of goods between different industries is 1 yuan, and the pricing of goods is based on a 10% profit. The goods produced by each industry and producers are randomly traded, forming a trading loop in the market until all credit debt settlement transactions stop, and everyone's debt is fully repaid to obtain the final M2 monetary amount and the amount of wealth each person has.

Golang实现一个货币信贷和市场经济活动场景，基础参数：基础货币M1发行总量10000元，1家银行模拟银行存款准备金率10%，忽略税费问题，一个模拟市场1000个人和20个行业财富为0，人员平均分配，其中各行业头部10%的人可以申请到贷款，可用信贷资金直至全部放出至不可贷，市场货币充分流通，各行业利润率相等；各行业之间商品初始1块钱，商品定价以10%利润定价，各行业生产的商品和生产者之间随机交易，市场形成交易闭环，直至所有信贷债务清偿交易停止，每个人的债务全部清偿，获取最终的M2货币量和每个人拥有财富金额。（最终M2总量不变）

**注意事项：以上预设条件仅供简单模拟测试，禁止用于专业测试和商业参考！**


### test-demo

```
最终的M2货币量：10000
第 1 个人的实际货币总量：3
第 2 个人的实际货币总量：11
第 3 个人的实际货币总量：10
第 4 个人的实际货币总量：17
第 5 个人的实际货币总量：4
第 6 个人的实际货币总量：6
第 7 个人的实际货币总量：14
第 8 个人的实际货币总量：10
第 9 个人的实际货币总量：17
第 10 个人的实际货币总量：13
第 11 个人的实际货币总量：6
第 12 个人的实际货币总量：11
第 13 个人的实际货币总量：3
第 14 个人的实际货币总量：14
第 15 个人的实际货币总量：17
第 16 个人的实际货币总量：6
第 17 个人的实际货币总量：3
第 18 个人的实际货币总量：12
第 19 个人的实际货币总量：17
第 20 个人的实际货币总量：6
第 21 个人的实际货币总量：4
第 22 个人的实际货币总量：3
第 23 个人的实际货币总量：18
第 24 个人的实际货币总量：9
第 25 个人的实际货币总量：10
第 26 个人的实际货币总量：8
第 27 个人的实际货币总量：11
第 28 个人的实际货币总量：4
第 29 个人的实际货币总量：9
第 30 个人的实际货币总量：9
第 31 个人的实际货币总量：11
第 32 个人的实际货币总量：9
第 33 个人的实际货币总量：10
第 34 个人的实际货币总量：17
第 35 个人的实际货币总量：7
第 36 个人的实际货币总量：1
第 37 个人的实际货币总量：13
第 38 个人的实际货币总量：10
第 39 个人的实际货币总量：13
第 40 个人的实际货币总量：10
第 41 个人的实际货币总量：17
第 42 个人的实际货币总量：11
第 43 个人的实际货币总量：14
第 44 个人的实际货币总量：9
第 45 个人的实际货币总量：12
第 46 个人的实际货币总量：11
第 47 个人的实际货币总量：12
第 48 个人的实际货币总量：5
第 49 个人的实际货币总量：13
第 50 个人的实际货币总量：18
第 51 个人的实际货币总量：15
第 52 个人的实际货币总量：13
第 53 个人的实际货币总量：17
第 54 个人的实际货币总量：5
第 55 个人的实际货币总量：8
第 56 个人的实际货币总量：16
第 57 个人的实际货币总量：11
第 58 个人的实际货币总量：9
第 59 个人的实际货币总量：16
第 60 个人的实际货币总量：9
第 61 个人的实际货币总量：9
第 62 个人的实际货币总量：6
第 63 个人的实际货币总量：13
第 64 个人的实际货币总量：16
第 65 个人的实际货币总量：2
第 66 个人的实际货币总量：12
第 67 个人的实际货币总量：15
第 68 个人的实际货币总量：10
第 69 个人的实际货币总量：13
第 70 个人的实际货币总量：11
第 71 个人的实际货币总量：17
第 72 个人的实际货币总量：7
第 73 个人的实际货币总量：8
第 74 个人的实际货币总量：11
第 75 个人的实际货币总量：6
第 76 个人的实际货币总量：8
第 77 个人的实际货币总量：8
第 78 个人的实际货币总量：9
第 79 个人的实际货币总量：16
第 80 个人的实际货币总量：7
第 81 个人的实际货币总量：7
第 82 个人的实际货币总量：17
第 83 个人的实际货币总量：9
第 84 个人的实际货币总量：13
第 85 个人的实际货币总量：5
第 86 个人的实际货币总量：10
第 87 个人的实际货币总量：13
第 88 个人的实际货币总量：8
第 89 个人的实际货币总量：7
第 90 个人的实际货币总量：14
第 91 个人的实际货币总量：14
第 92 个人的实际货币总量：8
第 93 个人的实际货币总量：12
第 94 个人的实际货币总量：14
第 95 个人的实际货币总量：4
第 96 个人的实际货币总量：14
第 97 个人的实际货币总量：9
第 98 个人的实际货币总量：14
第 99 个人的实际货币总量：8
第 100 个人的实际货币总量：15
... ...

```
