using JetBrains.Annotations;
using Unii.Error;

namespace Unii.Scan;

[PublicAPI]
public class Scanner
{
    private static readonly Dictionary<string, TokenType> Keywords = new()
    {
        { "true", TokenType.True },
        { "false", TokenType.False },
        { "null", TokenType.Null },
        { "var", TokenType.Var },
        { "const", TokenType.Const },
    };

    private readonly string _src;
    private readonly ErrorHandler _errorHandler;
    private readonly List<Token> _tokens = new();

    private int _start;
    private int _current;
    private int _line = 1;

    public Scanner(string src, ErrorHandler errorHandler)
    {
        _src = src;
        _errorHandler = errorHandler;
    }
    
    private bool IsAtEnd()
    {
        return _current >= _src.Length;
    }
    
    private char Advance()
    {
        return _src[_current++];
    }
    
    private bool Match(char expected)
    {
        if (IsAtEnd()) return false;
        if (_src[_current] != expected) return false;
        
        _current++;
        return true;
    }
    
    private char Peek(int offset = 0)
    {
        return _current + offset >= _src.Length ? '\0' : _src[_current + offset];
    }
    
    private Token AddToken(TokenType type)
    {
        var token = new Token(type, _src[_start.._current], _line, _start+1);
        _tokens.Add(token);
        return token;
    }
    
    private void ScanString()
    {
        while (Peek() != '"' && !IsAtEnd() && Peek() != '\n')
        {
            if (Peek() == '\\' && Peek(1) == '"')
                Advance();
            Advance();
        }

        if (IsAtEnd())
        {
            _errorHandler.PushError(new CompileError("Unterminated string", _line, _start));
            return;
        }

        Advance();
        AddToken(TokenType.StringLiteral);
    }

    private void ScanToken()
    {
        var c = Advance();

        switch (c)
        {
            case ' ':
                AddToken(TokenType.Space);
                break;
            case '\t':
                AddToken(TokenType.Tab);
                break;
            case '\n':
                AddToken(TokenType.NewLine);
                break;
            case '+':
                AddToken(Match('+') ? TokenType.Inc : Match('=') ? TokenType.AddAssign : TokenType.Add);
                break;
            case '-':
                AddToken(Match('-') ? TokenType.Dec : Match('=') ? TokenType.SubAssign : TokenType.Sub);
                break;
            case '*':
                if (Match('*'))
                {
                    AddToken(Match('=') ? TokenType.PowerAssign : TokenType.Power);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.MulAssign : TokenType.Mul);
                }
                break;
            case '/':
                if (Match('/'))
                {
                    while (Peek() != '\n' && !IsAtEnd()) Advance();
                    AddToken(TokenType.SingleLineComment);
                }
                else if (Match('*'))
                {
                    while (Peek() != '*' && Peek(1) != '/' && !IsAtEnd())
                    {
                        if (Peek() == '\n') _line++;
                        Advance();
                    }
                    if (IsAtEnd())
                    {
                        _errorHandler.PushError(new CompileError("Unterminated comment", _line, _start));
                        return;
                    }
                    Advance();
                    Advance();
                    AddToken(TokenType.MultiLineComment);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.DivAssign : TokenType.Div);
                }
                break;
            case '%':
                AddToken(Match('=') ? TokenType.ModAssign : TokenType.Mod);
                break;
            case '(':
                AddToken(TokenType.OpenParen);
                break;
            case ')':
                AddToken(TokenType.CloseParen);
                break;
            case '[':
                AddToken(TokenType.OpenBracket);
                break;
            case ']':
                AddToken(TokenType.CloseBracket);
                break;
            case '{':
                AddToken(TokenType.OpenBrace);
                break;
            case '}':
                AddToken(TokenType.CloseBrace);
                break;
            case ':':
                AddToken(TokenType.Colon);
                break;
            case ';':
                AddToken(TokenType.SemiColon);
                break;
            case ',':
                AddToken(TokenType.Comma);
                break;
            case '.': 
                AddToken(TokenType.Dot);
                break;
            case '=':
                AddToken(Match('=') ? TokenType.Equal : TokenType.Assign);
                break;
            case '<':
                if (Match('<'))
                {
                    AddToken(Match('=') ? TokenType.BitwiseLeftShiftAssign : TokenType.BitwiseLeftShift);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.LessThanOrEqual : TokenType.LessThan);
                }
                break;
            case '>':
                if (Match('>'))
                {
                    AddToken(Match('=') ? TokenType.BitwiseRightShiftAssign : TokenType.BitwiseRightShift);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.GreaterThanOrEqual : TokenType.GreaterThan);
                }
                break;
            case '!':
                AddToken(Match('=') ? TokenType.NotEqual : TokenType.Not);
                break;
            case '&':
                if (Match('&'))
                {
                    AddToken(Match('=') ? TokenType.AndAssign : TokenType.And);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.BitwiseAndAssign : TokenType.BitwiseAnd);
                }
                break;
            case '|':
                if (Match('|'))
                {
                    AddToken(Match('=') ? TokenType.OrAssign : TokenType.Or);
                }
                else
                {
                    AddToken(Match('=') ? TokenType.BitwiseOrAssign : TokenType.BitwiseOr);
                }
                break;
            case '^':
                AddToken(Match('=') ? TokenType.BitwiseXorAssign : TokenType.BitwiseXor);
                break;
            case '~':
                AddToken(TokenType.BitwiseNot);
                break;
            case '"':
                ScanString();
                break;
            case '\'':
                if (Peek() == '\\' && Peek(1) == '\'')
                    Advance();
                Advance();
                if (Peek() != '\'')
                {
                    _errorHandler.PushError(new CompileError("Expected closing single quote", _line, _start));
                    return;
                }
                Advance();
                AddToken(TokenType.CharacterLiteral);
                break;
            default:
                if (char.IsDigit(c))
                {
                    var isFloat = false;
                    while (char.IsDigit(Peek())) Advance();
                    if (Peek() == '.' && char.IsDigit(Peek(1)))
                    {
                        isFloat = true;
                        Advance();
                        while (char.IsDigit(Peek())) Advance();
                    }
                    AddToken(isFloat ? TokenType.FloatLiteral : TokenType.IntegerLiteral);
                }
                else if (char.IsLetter(c) || c == '_')
                {
                    while (char.IsLetterOrDigit(Peek()) || Peek() == '_') Advance();
                    var text = _src[_start.._current];
                    AddToken(Keywords.TryGetValue(text, out var keyword) ? keyword : TokenType.Identifier);
                }
                else _errorHandler.PushError(new CompileError($"Unexpected character '{c}'", _line, _start));
                break;
        }
    }

    public List<Token> ScanTokens()
    {   
        while (!IsAtEnd())
        {
            _start = _current;
            ScanToken();
        }

        _tokens.Add(new Token(TokenType.Eof, "", _line, _start+1));

        return _tokens;
    }
}