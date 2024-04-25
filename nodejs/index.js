import express from 'express'
import { createConnection } from 'mysql2/promise'
import { randomUUID } from 'node:crypto'

const app = express();

const connection = await createConnection({
  host: 'mysql',
  user: 'root',
  password: 'root',
  database: 'bench'
});

app.use(express.json())

app.get('/hello', async (req, res) => {
  res.send('Hello World')
})

app.post('/users', async (req, res) => {
  try {
    const id = randomUUID()
    const { name } = req.body
    await connection.execute('INSERT INTO users (id, name) values (?, ?)', [id, name])
    res.sendStatus(201)
  } catch (error) {
    console.log('error :', error);
    res.status(500).send('Erro ao criar usuario')
  }
})

app.listen(8080, () => {
  console.log('Server On')
})


