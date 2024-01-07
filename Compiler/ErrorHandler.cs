namespace Compiler;

public class ErrorCompiler
{
    public readonly string Message;
    public readonly int Line;
    public readonly int Column;
    
    public ErrorCompiler(string message, int line, int column)
    {
        Message = message;
        Line = line;
        Column = column;
    }
}

public class ErrorHandler
{
    private readonly List<ErrorCompiler> _errors = new();
    
    public int PushError(ErrorCompiler error)
    {
        _errors.Add(error);
        return _errors.Count - 1;
    }
    
    public ErrorCompiler PopError()
    {
        var error = _errors[^1];
        _errors.RemoveAt(_errors.Count - 1);
        return error;
    }
    
    public ErrorCompiler GetError(int index)
    {
        return _errors[index];
    }
    
    public int GetErrorCount()
    {
        return _errors.Count;
    }

    public void Report()
    {
        foreach (var error in _errors)
        {
            Console.WriteLine($"error ({error.Line}:{error.Column}): {error.Message}");
        }
    }
}