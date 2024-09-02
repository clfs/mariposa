mod core;

fn main() {
    let c = core::Color::White;

    println!("Current color: {:?}", c);
    println!("Other color: {:?}", c.other());
}
