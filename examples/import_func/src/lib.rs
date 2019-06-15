use std::ffi::{CString};
use std::os::raw::{c_char};

extern "C" {
    fn sum(x: i32, y: i32) -> i32;
    fn repeat(pointer: *const u8, length: u32, count: i32, output: *const u8, outLen: u32) -> i32;
}

#[no_mangle]
pub extern "C" fn add1(x: i32, y: i32) -> *mut c_char {
    let msg = "fool ";
    let mut response = b"                                                                                                                                                               ".to_vec();

    unsafe { 
        let cnt = sum(x, y) + 1;
        let _len = repeat(msg.as_ptr(), msg.len() as u32, cnt, response.as_ptr(), response.len() as u32);
        return CString::from_vec_unchecked(response).into_raw();
    }
}
