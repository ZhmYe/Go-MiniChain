package config

/**
 * 配置文件
 */

// Config 该类为配置类，主要有两个字段：
// difficulty: 挖矿的难度值，即规定了新的区块的哈希值至少以几个0开头才满足难度条件
// maxTransactionCount: 交易池大小；TransactionProducer需要随机生成交易，放入交易池中，直至达到该大小
type Config struct {
	difficulty          int
	maxTransactionCount int
	nbAccount           int
	initAmout           int
}

func (c *Config) GetDifficulty() int {
	return c.difficulty
}
func (c *Config) GetMaxTransactionCount() int {
	return c.maxTransactionCount
}
func (c *Config) GetAccountNumber() int {
	return c.nbAccount
}
func (c *Config) GetInitAmount() int {
	return c.initAmout
}

var MiniChainConfig = Config{
	difficulty:          2,
	maxTransactionCount: 64,
	nbAccount:           100,
	initAmout:           10000,
}
