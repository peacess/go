# see
[MCP server for Azure Cosmos DB using the Go SDK](https://github.com/abhirockzz/mcp_cosmosdb_go/tree/main)  

# 对话方式
## 直接对话，
## 每次对话时，都加上历史记录（历史记录就是上下文）
## 微调，给模型提供训练数据来创建微调模型，然后使用微调后的模型，这样对话会使用到我们自己提供的训练数据，不用每次对话都带上，以减少token的使用
## 插件，向chatgpt注册插件，由openai自己确认调用那个插件来提供数据或执行操作
## MCP，无需向chatgpt注册，是开放协议，多家平台都可以使用。需要ai host支持 MCP协议
