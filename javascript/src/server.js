const express = require('express');
const { getItems, createItem, getItem, updateItem, deleteItem } = require('./handlers');

const app = express();
app.use(express.json());

app.get('/ping', (req, res) => {
  res.send('pong');
})

app.get('/items', getItems);
app.post('/items', createItem);
app.get('/items/:id', getItem);
app.put('/items/:id', updateItem);
app.delete('/items/:id', deleteItem);

const PORT = 8080;

app.listen(PORT, () => {
  console.log(`App listening at http://localhost:${PORT}`);
});
