mod core;

fn main() {
    let c = core::Color::White;

    println!("Current color: {:?}", c);
    println!("Other color: {:?}", c.other());

    let s = core::Square::new(core::File::E, core::Rank::Second);

    println!("Current square: {:?}", s);
    println!("Current file: {:?}", s.file());
    println!("Current rank: {:?}", s.rank());
}
