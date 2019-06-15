use std::ffi::{CString};
use std::mem;
use std::os::raw::{c_char, c_void};

extern "C" {
    fn sum(x: i32, y: i32) -> i32;
    fn repeat(pointer: *const u8, length: u32, count: i32, output: *const u8, outLen: u32) -> i32;
}

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

#[no_mangle]
pub extern "C" fn add1(x: i32, y: i32) -> *mut c_char {
    let msg = "fool ";
    let mut response = vec![0; 2048];

    unsafe { 
        let cnt = sum(x, y) + 1;
        let _len = repeat(msg.as_ptr(), msg.len() as u32, cnt, response.as_ptr(), response.len() as u32);
        return CString::from_vec_unchecked(response).into_raw();
    }
}
