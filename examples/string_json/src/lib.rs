extern crate serde;
extern crate serde_json_core;

use std::ffi::{CStr, CString};
use std::mem;
use std::os::raw::{c_char, c_void};

use serde::{Deserialize, Serialize};
use serde_json_core::de::from_slice;

#[no_mangle]
pub extern "C" fn allocate(size: usize) -> *mut c_void {
    let mut buffer = Vec::with_capacity(size);
    let pointer = buffer.as_mut_ptr();
    mem::forget(buffer);

    pointer as *mut c_void
}

#[no_mangle]
pub extern "C" fn deallocate(pointer: *mut c_void, capacity: usize) {
    unsafe {
        let _ = Vec::from_raw_parts(pointer, 0, capacity);
    }
}

#[derive(Serialize, Deserialize)]
struct MyDumbStruct<'a> {
    number: i32,
    message: &'a str
}

#[no_mangle]
pub extern "C" fn greet(subject: *mut c_char) -> *mut c_char {
    let subject = unsafe { CStr::from_ptr(subject).to_bytes().to_vec() };

    let my_dumb_struct: MyDumbStruct = from_slice(&subject).unwrap();
    let mut output = b"Hello, ".to_vec();

    for _ in 0..my_dumb_struct.number {
        output.extend(my_dumb_struct.message.bytes());
    }

    output.extend(&[b'!']);

    unsafe { CString::from_vec_unchecked(output) }.into_raw()
}
