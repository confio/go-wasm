extern "C" {
    fn sum(x: i32, y: i32) -> i32;
    fn repeat(pointer: *const u8, length: u32, count: i32) -> i32;
}

#[no_mangle]
pub extern "C" fn add1(x: i32, y: i32) -> i32 {
    let msg = "fool ";

    unsafe { 
        let cnt = sum(x, y) + 1;
        let len = repeat(msg.as_ptr(), msg.len() as u32, cnt);
        return len;
    }
}
