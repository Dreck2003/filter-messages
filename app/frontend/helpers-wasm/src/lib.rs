mod helpers;
mod utils;

use wasm_bindgen::prelude::*;

use helpers::strings::replace_insensitive;

// When the `wee_alloc` feature is enabled, use `wee_alloc` as the global
// allocator.
#[cfg(feature = "wee_alloc")]
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;

// #[wasm_bindgen]
// extern "C" {
//     fn alert(s: &str);
// }

#[wasm_bindgen]
pub fn parse_email_to_html(text: String, expr: String, tag: Option<String>) -> String {
    let mut parse_text = replace_insensitive(&text, &expr.to_lowercase(), tag.as_deref());
    parse_text = parse_text.replace("<", "&lt;");
    parse_text = parse_text.replace(">", "&gt;");
    return parse_text;
}
