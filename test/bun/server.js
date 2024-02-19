import { serve } from "bun";

const users = new Map(); // Einfache In-Memory-Datenstruktur

function getUsers() {
  return Array.from(users.values());
}

function getUser(id) {
  return users.get(id) || null;
}

function addUser(user) {
  users.set(user.id, user);
  return user;
}

function deleteUser(id) {
  return users.delete(id);
}

serve({
  port: 8888,
  fetch(request) {
    const url = new URL(request.url);
    const id = url.pathname.split('/')[2];

    switch (url.pathname) {
      case "/users":
        if (request.method === "GET") {
          return new Response(JSON.stringify(getUsers()), {
            headers: {
              "Content-Type": "application/json",
            },
          });
        } else if (request.method === "POST") {
          return request.json().then((user) => {
            addUser(user);
            return new Response(JSON.stringify(user), {
              headers: {
                "Content-Type": "application/json",
              },
              status: 201,
            });
          });
        }
        break;

      case `/users/${id}`:
        if (request.method === "GET") {
          const user = getUser(id);
          if (!user) {
            return new Response("User not found", { status: 404 });
          }
          return new Response(JSON.stringify(user), {
            headers: {
              "Content-Type": "application/json",
            },
          });
        } else if (request.method === "DELETE") {
          const deleted = deleteUser(id);
          if (!deleted) {
            return new Response("User not found", { status: 404 });
          }
          return new Response("User deleted", { status: 200 });
        }
        break;

      default:
        return new Response("Not Found", { status: 404 });
    }

    return new Response(null, { status: 405 });
  },
});

