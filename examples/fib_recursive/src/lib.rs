extern crate serde;
extern crate serde_json_core;

use serde::{Deserialize, Serialize};
use serde_json_core::de::from_str;

#[derive(Serialize, Deserialize)]
struct MyDumbStruct {
	number: i32,
}

#[no_mangle]
pub extern "C" fn fib(n: i32) -> i32 {
	if n == 1 || n == 2 {
		1
	} else {
		fib(n - 1) + fib(n - 2)
	}
}

#[no_mangle]
pub extern "C" fn app_main() -> i32 {
	let str = r#"{ "number": 35 }"#;

	let my_dumb_struct: MyDumbStruct = from_str(&str).unwrap();

	fib(my_dumb_struct.number)
}
