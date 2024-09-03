//! Implements basic chess functionality.

use std::convert::TryFrom;

#[derive(Debug, PartialEq)]
/// A color on the chess board.
pub enum Color {
    White,
    Black,
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

/// A file on the board.
#[derive(Debug, PartialEq)]
pub enum File {
    A,
    B,
    C,
    D,
    E,
    F,
    G,
    H,
}

impl TryFrom<u8> for File {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(File::A),
            1 => Ok(File::B),
            2 => Ok(File::C),
            3 => Ok(File::D),
            4 => Ok(File::E),
            5 => Ok(File::F),
            6 => Ok(File::G),
            7 => Ok(File::H),
            _ => Err(value),
        }
    }
}

/// A rank on the board.
#[derive(Debug, PartialEq)]
pub enum Rank {
    First,
    Second,
    Third,
    Fourth,
    Fifth,
    Sixth,
    Seventh,
    Eighth,
}

impl TryFrom<u8> for Rank {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Rank::First),
            1 => Ok(Rank::Second),
            2 => Ok(Rank::Third),
            3 => Ok(Rank::Fourth),
            4 => Ok(Rank::Fifth),
            5 => Ok(Rank::Sixth),
            6 => Ok(Rank::Seventh),
            7 => Ok(Rank::Eighth),
            _ => Err(value),
        }
    }
}

/// A square on the board.
#[derive(Clone, Copy, Debug, PartialEq)]
pub enum Square {
    A1,
    B1,
    C1,
    D1,
    E1,
    F1,
    G1,
    H1,
    A2,
    B2,
    C2,
    D2,
    E2,
    F2,
    G2,
    H2,
    A3,
    B3,
    C3,
    D3,
    E3,
    F3,
    G3,
    H3,
    A4,
    B4,
    C4,
    D4,
    E4,
    F4,
    G4,
    H4,
    A5,
    B5,
    C5,
    D5,
    E5,
    F5,
    G5,
    H5,
    A6,
    B6,
    C6,
    D6,
    E6,
    F6,
    G6,
    H6,
    A7,
    B7,
    C7,
    D7,
    E7,
    F7,
    G7,
    H7,
    A8,
    B8,
    C8,
    D8,
    E8,
    F8,
    G8,
    H8,
}

impl TryFrom<u8> for Square {
    type Error = u8;

    fn try_from(value: u8) -> Result<Self, Self::Error> {
        match value {
            0 => Ok(Self::A1),
            1 => Ok(Self::B1),
            2 => Ok(Self::C1),
            3 => Ok(Self::D1),
            4 => Ok(Self::E1),
            5 => Ok(Self::F1),
            6 => Ok(Self::G1),
            7 => Ok(Self::H1),
            8 => Ok(Self::A2),
            9 => Ok(Self::B2),
            10 => Ok(Self::C2),
            11 => Ok(Self::D2),
            12 => Ok(Self::E2),
            13 => Ok(Self::F2),
            14 => Ok(Self::G2),
            15 => Ok(Self::H2),
            16 => Ok(Self::A3),
            17 => Ok(Self::B3),
            18 => Ok(Self::C3),
            19 => Ok(Self::D3),
            20 => Ok(Self::E3),
            21 => Ok(Self::F3),
            22 => Ok(Self::G3),
            23 => Ok(Self::H3),
            24 => Ok(Self::A4),
            25 => Ok(Self::B4),
            26 => Ok(Self::C4),
            27 => Ok(Self::D4),
            28 => Ok(Self::E4),
            29 => Ok(Self::F4),
            30 => Ok(Self::G4),
            31 => Ok(Self::H4),
            32 => Ok(Self::A5),
            33 => Ok(Self::B5),
            34 => Ok(Self::C5),
            35 => Ok(Self::D5),
            36 => Ok(Self::E5),
            37 => Ok(Self::F5),
            38 => Ok(Self::G5),
            39 => Ok(Self::H5),
            40 => Ok(Self::A6),
            41 => Ok(Self::B6),
            42 => Ok(Self::C6),
            43 => Ok(Self::D6),
            44 => Ok(Self::E6),
            45 => Ok(Self::F6),
            46 => Ok(Self::G6),
            47 => Ok(Self::H6),
            48 => Ok(Self::A7),
            49 => Ok(Self::B7),
            50 => Ok(Self::C7),
            51 => Ok(Self::D7),
            52 => Ok(Self::E7),
            53 => Ok(Self::F7),
            54 => Ok(Self::G7),
            55 => Ok(Self::H7),
            56 => Ok(Self::A8),
            57 => Ok(Self::B8),
            58 => Ok(Self::C8),
            59 => Ok(Self::D8),
            60 => Ok(Self::E8),
            61 => Ok(Self::F8),
            62 => Ok(Self::G8),
            63 => Ok(Self::H8),
            _ => Err(value),
        }
    }
}

impl Square {
    /// Returns a new square on the given file and rank.
    pub fn new(f: File, r: Rank) -> Square {
        let ff = f as u8;
        let rr = r as u8;
        Square::try_from(rr * 8 + ff).unwrap()
    }

    /// Returns the file the square is on.
    pub fn file(self) -> File {
        let n = self as u8;
        let f = n % 8;
        File::try_from(f).unwrap()
    }

    /// Returns the rank the square is on.
    pub fn rank(self) -> Rank {
        let n = self as u8;
        let r = n / 8;
        Rank::try_from(r).unwrap()
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

    #[test]
    fn square_new() {
        let f = core::File::F;
        let r = core::Rank::Sixth;
        let s = core::Square::new(f, r);
        assert_eq!(s, core::Square::F6);
    }

    #[test]
    fn square_file() {
        let f = core::Square::B3.file();
        assert_eq!(f, core::File::B);
    }

    #[test]
    fn square_rank() {
        let r = core::Square::H4.rank();
        assert_eq!(r, core::Rank::Fourth);
    }
}
