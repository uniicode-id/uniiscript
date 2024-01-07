namespace Compiler.Scanning;

public enum TokenType
{
    OpenParen,
    CloseParen,
    OpenBracket,
    CloseBracket,
    OpenBrace,
    CloseBrace,
    
    Colon,
    SemiColon,
    Comma,
    Dot,
    
    // Comment
    SingleLineComment,
    MultiLineComment,
    
    // White Space
    Space,
    Tab,
    NewLine,

    // Arithmetic Operator
    Add,
    Sub,
    Mul,
    Div,
    Mod,
    Power,
    Inc,
    Dec,

    // Assignment Operator
    Assign,
    AddAssign,
    SubAssign,
    MulAssign,
    DivAssign,
    ModAssign,
    PowerAssign,
    BitwiseAndAssign,
    BitwiseOrAssign,
    BitwiseXorAssign,
    BitwiseLeftShiftAssign,
    BitwiseRightShiftAssign,
    AndAssign,
    OrAssign,

    // Bitwise Operator
    BitwiseAnd,
    BitwiseOr,
    BitwiseXor,
    BitwiseNot,
    BitwiseLeftShift,
    BitwiseRightShift,

    // Comparator Operator
    Equal,
    NotEqual,
    LessThan,
    LessThanOrEqual,
    GreaterThan,
    GreaterThanOrEqual,
    
    // Logical Operator
    And,
    Or,
    Not,
    
    // Identifier
    Identifier,
    
    // Literal
    StringLiteral,
    IntegerLiteral,
    FloatLiteral,
    CharacterLiteral,
    
    // Keyword
    Var,
    Const,
    True,
    False,
    Null,
    
    Eof
}