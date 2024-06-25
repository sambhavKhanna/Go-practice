const express = require('express')
const app = express()
const port = 3000
let num = 0
app.use(express.json())

app.get('/', (req, res) => {
    num++
    console.log("REQUEST MADE: ", num)
    res.status(200).json({ message: 'working' })
})

app.post('/', (req, res) => {
    num++
    console.log("REQUEST MADE: ", num)
    res.status(200).json(req.body)
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})