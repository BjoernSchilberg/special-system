


## cURL

```shell
curl -X POST -H "Content-Type: application/json" -d '{"id":"3", "name":"Alice Smith"}' http://localhost:8080/users
```

```shell
curl -X DELETE http://localhost:8888/users/3
```


## Benchmark HTTP-Performance

```shell
siege -c10 -t1M http://127.0.0.1:8888/users
```

Go

```text

{	"transactions":			     1459371,
	"availability":			      100.00,
	"elapsed_time":			       59.08,
	"data_transferred":		       83.51,
	"response_time":		        0.00,
	"transaction_rate":		    24701.61,
	"throughput":			        1.41,
	"concurrency":			        9.22,
	"successful_transactions":   1459371,
	"failed_transactions":		       0,
	"longest_transaction":		    0.02,
	"shortest_transaction":		    0.00
}
```

Rust

```shell
cargo run
```
```text
{	"transactions":			      803379,
	"availability":			      100.00,
	"elapsed_time":			       59.67,
	"data_transferred":		        1.53,
	"response_time":		        0.00,
	"transaction_rate":		    13463.70,
	"throughput":			        0.03,
	"concurrency":			        9.43,
	"successful_transactions":	      803379,
	"failed_transactions":		           0,
	"longest_transaction":		        0.02,
	"shortest_transaction":		        0.00
}
```



Python/Flask
in `test/test.py`
```text
{
   "transactions":			       94275,
	"availability":			      100.00,
	"elapsed_time":			       59.16,
	"data_transferred":		        8.72,
	"response_time":		        0.01,
	"transaction_rate":		     1593.56,
	"throughput":			        0.15,
	"concurrency":			        9.95,
	"successful_transactions":	   94275,
	"failed_transactions":		       0,
	"longest_transaction":		    0.03,
	"shortest_transaction":		    0.00
}
```

Odin/Zig

Odin and Zig do not have a direct standard library or widely used third-party
libraries specifically for creating HTTP servers, as is the case in languages
such as Rust with warp or Go with its net/http package.
