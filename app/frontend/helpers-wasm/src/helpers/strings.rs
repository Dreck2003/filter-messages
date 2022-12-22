pub fn replace_insensitive(text: &str, expr: &str, tag: Option<&str>) -> String {
    match tag {
        Some(html_tag) => {
            return matches_insensitive(text, expr, |s| {
                return format!("<{}>{}</{}>", html_tag, s, html_tag);
            });
        }
        None => {
            return matches_insensitive(text, expr, |s| {
                return format!("<mark>{}</mark>", s);
            });
        }
    }
}

pub fn matches_insensitive<F>(text: &str, expr: &str, change: F) -> String
where
    F: Fn(&str) -> String,
{
    if text.is_empty() {
        return "".to_string();
    }
    let mut index = 0;
    let mut mutate_in = 0;
    let max_size = expr.len();
    let str_length = text.len();
    let lower_expr = expr.to_lowercase();
    let first_letter = expr.chars().next().unwrap();

    let mut string_state = String::new();
    let mut start_with_first_letter = true;
    loop {
        mutate_in += 1;
        if mutate_in >= 5000 {
            // This is just to avoid the infinity loop
            return string_state;
        }
        if index + max_size > str_length {
            break;
        }
        let hold_str = text.get(index..(index + max_size));
        match hold_str {
            Some(hold_match) => {
                let compare = hold_match.to_lowercase();
                if compare == lower_expr {
                    // TODO!: match!
                    index += max_size;
                    start_with_first_letter = true;
                    string_state.push_str(&change(hold_match)); // concat the str
                    continue;
                }
                let init = if !start_with_first_letter { 1 } else { 0 };
                // TODO: not match
                let mut find_first_letter_match = false;
                if !start_with_first_letter {
                    index += 1;
                    string_state.push_str(&hold_match[0..1])
                }
                for letter in hold_match[init..].chars() {
                    if letter.to_lowercase().next().unwrap() == first_letter {
                        find_first_letter_match = true;
                        break;
                    }
                    index += 1;
                    string_state.push(letter);
                }
                start_with_first_letter = if find_first_letter_match { false } else { true };
                continue;
            }
            None => break,
        }
    }
    string_state.push_str(&text[index..]);
    string_state
}

#[cfg(test)]
mod test_strings {
    use super::*;

    #[test]
    fn test_replace_with_custom_tag() {
        let mark = "high";
        let text_hightligthed = "<high>will</high>";

        let res = replace_insensitive(
            "This is text that will replace, and other will word",
            "will",
            Some(mark),
        );
        let parse_res = format!(
            "This is text that {} replace, and other {} word",
            text_hightligthed, text_hightligthed
        );

        assert_eq!(parse_res, res);
    }
    #[test]
    fn test_replace_with_default() {
        let res = replace_insensitive(
            "Other Text with more text word and other text :)",
            "text",
            None,
        );
        let text_hightligthed = "<mark>text</mark>";
        let parse_res = format!(
            "Other {} with more {} word and other {} :)",
            "<mark>Text</mark>", text_hightligthed, text_hightligthed
        );

        assert_eq!(parse_res, res)
    }

    #[test]
    fn test_matches_insensitive() {
        let text = "Will is changed and this will too";
        let res = matches_insensitive(text, "will", |s| format!("<mark>{}</mark>", s));
        println!("{}", res);

        assert_eq!(
            "<mark>Will</mark> is changed and this <mark>will</mark> too",
            res
        );
    }
    #[test]
    fn test_matches_insensitive_empty() {
        let text = "";
        let res = matches_insensitive(text, "", |s| format!("<mark>{}</mark>", s));
        assert_eq!("", res);
    }
}
