import { ApolloServer } from "apollo-server"
import { ApolloGateway } from "@apollo/gateway"

const gateway = new ApolloGateway({
  serviceList: [{ name: "serviceA", url: "http://localhost:4001/query" }],
})

const server = new ApolloServer({
  gateway,
  subscriptions: false,
})

server.listen().then(({ url }) => {
  console.log(`Server is ready at ${url}`)
})
