use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use warp::Filter;

#[derive(Debug, Clone, Deserialize, Serialize)]
struct User {
    id: String,
    name: String,
}

type Users = Arc<Mutex<HashMap<String, User>>>;

async fn get_users(users: Users) -> Result<impl warp::Reply, warp::Rejection> {
    let users = users.lock().unwrap();
    let users: Vec<User> = users.values().cloned().collect();
    Ok(warp::reply::json(&users))
}

async fn get_user(id: String, users: Users) -> Result<impl warp::Reply, warp::Rejection> {
    let users = users.lock().unwrap();
    if let Some(user) = users.get(&id) {
        Ok(warp::reply::json(user))
    } else {
        Err(warp::reject::not_found())
    }
}

async fn create_user(new_user: User, users: Users) -> Result<impl warp::Reply, warp::Rejection> {
    let mut users = users.lock().unwrap();
    users.insert(new_user.id.clone(), new_user.clone());
    Ok(warp::reply::json(&new_user))
}

async fn delete_user(id: String, users: Users) -> Result<impl warp::Reply, warp::Rejection> {
    let mut users = users.lock().unwrap();
    users.remove(&id);
    Ok(warp::reply::with_status(
        "User deleted",
        warp::http::StatusCode::OK,
    ))
}

#[tokio::main]
async fn main() {
    let users = Users::default();
    let users_filter = warp::any().map(move || users.clone());

    let get_users_route = warp::path("users")
        .and(warp::get())
        .and(users_filter.clone())
        .and_then(get_users);

    let get_user_route = warp::path!("users" / String)
        .and(warp::get())
        .and(users_filter.clone())
        .and_then(get_user);

    let create_user_route = warp::path("users")
        .and(warp::post())
        .and(warp::body::json())
        .and(users_filter.clone())
        .and_then(create_user);

    let delete_user_route = warp::path!("users" / String)
        .and(warp::delete())
        .and(users_filter.clone())
        .and_then(delete_user);

    let routes = get_users_route
        .or(get_user_route)
        .or(create_user_route)
        .or(delete_user_route)
        .with(warp::cors().allow_any_origin());

    warp::serve(routes).run(([127, 0, 0, 1], 8888)).await;
}
