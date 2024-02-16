from flask import Flask, jsonify, request, abort

app = Flask(__name__)

# Einfache In-Memory-Datenbank für Demonstrationszwecke
users = [
    {"id": "1", "name": "John Doe"},
    {"id": "2", "name": "Jane Doe"},
]

# Funktionen zur Handhabung der Anfragen
@app.route('/users', methods=['GET'])
def get_users():
    return jsonify(users)

@app.route('/users/<string:id>', methods=['GET'])
def user_by_id(id):
    user = next((user for user in users if user['id'] == id), None)
    if user is not None:
        return jsonify(user)
    else:
        abort(404)

@app.route('/users', methods=['POST'])
def create_user():
    if not request.json or not 'name' in request.json or not 'id' in request.json:
        abort(400)
    user = {
        'id': request.json['id'],
        'name': request.json['name']
    }
    users.append(user)
    return jsonify(user), 201

@app.route('/users/<string:id>', methods=['DELETE'])
def delete_user(id):
    global users
    users = [user for user in users if user['id'] != id]
    return jsonify({'result': True})

if __name__ == '__main__':
    app.run(debug=True, port=8888)

