const Web3 = require('web3')
const rpcURL = 'https://cloudflare-eth.com' // Your RCkP URL goes here
const web3 = new Web3(rpcURL)
const address = '0xa5b5E8754e87E9e407F548128855468db9747126' // Your account address goes here
web3.eth.getBalance(address, (err, wei) => { balance = web3.utils.fromWei(wei, 'ether') }).then(console.log(wei))
