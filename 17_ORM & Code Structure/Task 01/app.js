const express = require('express');
const bodyParser = require('body-parser');
const packageRoutes = require('./routes/packageRoutes');
const sequelize = require('./config/database');

const app = express();
app.use(bodyParser.json());
app.use(packageRoutes);

sequelize.sync()
  .then(() => {
    console.log('Database synced');
  })
  .catch(err => {
    console.error('Error syncing database:', err);
  });

module.exports = app;