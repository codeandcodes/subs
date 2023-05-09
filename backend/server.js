const express = require('express');
const cors = require('cors');
const fs = require('fs');
const yaml = require('js-yaml');
const { createSquareClient, getCatalogItems, createSubscription, getLocations } = require('./square.js');
const { Client, Environment } = require('square');
const protobuf = require('google-protobuf');
const { Subscription } = require('./subscription_pb');

const app = express();

// Config

// Read Square API credentials from YAML or properties file
const configFile = process.argv[2];
const squareConfig = yaml.load(fs.readFileSync(configFile, 'utf8'));

// Square API credentials
const squareAccessToken = squareConfig.accessToken;
const squareEnvironment = squareConfig.environment === 'production' ? Environment.Production : Environment.Sandbox;

const squareClient = createSquareClient(squareAccessToken, squareEnvironment);


// Middleware
app.use(cors());
app.use(express.json());

// Routes
app.get('/', (req, res) => {
  res.send('Welcome to the backend!');
});

app.get('/location', (req, res) => {
    getLocations(squareClient);

    // Create a new subscription object
    const subscription = new Subscription();

    // Set some example fields
    subscription.setId('123');
    subscription.setAmount(1000);
    subscription.setDescription('Example subscription');
    

    // Log the fields
    console.log(subscription.getId());
    console.log(subscription.getAmount());
    console.log(subscription.getDescription());
    res.send('Welcome to the backend!');
        
        
});

// Start the server
const port = 3000;
app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
