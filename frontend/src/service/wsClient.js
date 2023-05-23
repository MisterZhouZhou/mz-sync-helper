const url = `ws://${window.location.hostname}:27149/ws`;
const wsClient = new WebSocket(url);

// socket客户端
class WsClient {
  constructor(client) {
    this.client = client
  }

  // 发送数据
  send(data) {
    this.client.send(JSON.stringify(data))
  }

  // 接收数据
  onMessage(fn) {
    this.client.onmessage = ({ data }) => {
      fn(JSON.parse(data))
    }
  }
}

const wsClientPromise = new Promise((resolve, reject) => {
  wsClient.onopen = () => {
    resolve(new WsClient(wsClient))
  }
  setTimeout(() => {
    reject(new Error('get ws connection timeout'))
  }, 10000)
})

export const getWsClient = () => wsClientPromise;