from flask import Flask, jsonify, request
import pymongo
import random
import time
import requests
import json

app = Flask(__name__)

# Connect to MongoDB and get the "products" collection
client = pymongo.MongoClient("mongodb://localhost:27017/")
db = client["mydatabase"]
products = db["products"]

# Define a function to generate random product data
def generate_product():
    return {
        "name": "Product " + str(random.randint(1, 1000)),
        "category": random.choice(["Electronics", "Clothing", "Books"]),
        "price": round(random.uniform(10, 100), 2),
        "stock": random.randint(1, 100)
    }

# Define a route to add new products
@app.route("/products", methods=["POST"])
def add_product():
    # Generate a random product and insert it into the "products" collection
    product = generate_product()
    products.insert_one(product)
    return json.dumps(product,  default=str)

# Define a route to retrieve products by category
@app.route("/products/<category>")
def get_products_by_category(category):
    # Find all products in the specified category and return them as a JSON array
    result = products.find({"category": category})
    return json.dumps(list(result),  default=str)

# Define a route to simulate user traffic
@app.route("/simulate-traffic")
def simulate_traffic():
    # Generate a random rate of requests (up to 10 per second) and send them to random routes
    routes = ["/products/Electronics", "/products/Clothing", "/products/Books", "/products"]
    rate = random.uniform(0, 10)
    print("Simulating traffic at a rate of {:.2f} requests per second".format(rate))
    start_time = time.time()
    while time.time() - start_time < 60:
        route = random.choice(routes)
        if route == "/products":
            # Add a new product
            product = generate_product()
            products.insert_one(product)
        else:
            # Get products by category
            result = requests.get("http://localhost:5000" + route)
        time.sleep(random.uniform(0, 1/rate))
    return "Simulated traffic for 60 seconds at a rate of {:.2f} requests per second".format(rate)

if __name__ == "__main__":
    app.run(debug=True)
