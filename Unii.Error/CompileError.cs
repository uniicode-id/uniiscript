using JetBrains.Annotations;

namespace Unii.Error;

[PublicAPI]
public class CompileError
{
    public readonly string Message;
    public readonly int Line;
    public readonly int Column;
    
    public CompileError(string message, int line, int column)
    {
        Message = message;
        Line = line;
        Column = column;
    }
}