using JetBrains.Annotations;

namespace Unii.Error;

[PublicAPI]
public class ErrorHandler
{
    private readonly List<CompileError> _errors = new();
    
    public int PushError(CompileError error)
    {
        _errors.Add(error);
        return _errors.Count - 1;
    }
    
    public CompileError PopError()
    {
        var error = _errors[^1];
        _errors.RemoveAt(_errors.Count - 1);
        return error;
    }
    
    public CompileError GetError(int index)
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