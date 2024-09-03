//! Implements basic chess functionality.

#[derive(Debug, PartialEq)]
/// A color on the chess board.
pub enum Color {
    White,
    Black
}

impl Color {
    /// Returns the other color.
    pub fn other(&self) -> Self {
        match self {
            Color::White => Color::Black,
            Color::Black => Color::White,
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::core;

    #[test]
    fn other_color() {
        let x = core::Color::White;

        let y = x.other();
        assert_eq!(y, core::Color::Black);

        let z = y.other();
        assert_eq!(z, core::Color::White);
    }
}
