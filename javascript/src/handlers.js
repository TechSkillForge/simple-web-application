let items = [];
let nextId = 1;

const getItems = (req, res) => {
    res.json(items);
};

const createItem = (req, res) => {
    const item = { id: nextId++, ...req.body };
    items.push(item);
    res.status(201).json(item);
};

const getItem = (req, res) => {
    const item = items.find(i => i.id === parseInt(req.params.id));
    if (item) {
        res.json(item);
    } else {
        res.status(404).send('Item not found');
    }
};

const updateItem = (req, res) => {
    const index = items.findIndex(i => i.id === parseInt(req.params.id));
    if (index !== -1) {
        // Update everything except the id
        const updatedItem = { ...items[index], ...req.body, id: items[index].id };
        items[index] = updatedItem;
        res.json(updatedItem);
    } else {
        res.status(404).send('Item not found');
    }
};

const deleteItem = (req, res) => {
    const index = items.findIndex(i => i.id === parseInt(req.params.id));
    if (index !== -1) {
        items = items.filter(i => i.id !== parseInt(req.params.id));
        res.status(200).send('Item deleted');
    } else {
        res.status(404).send('Item not found');
    }
};

module.exports = { getItems, createItem, getItem, updateItem, deleteItem };