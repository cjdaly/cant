use std::env;
use std::fs;
use minidom::Element;

fn main() {
  let args: Vec<String> = env::args().collect();
  let filename = &args[1];
  let text = fs::read_to_string(filename).expect("IO Error!");
  let root: Element = text.parse().unwrap();
  println!("{:#?}", root);
}

