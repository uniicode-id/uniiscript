using Compiler.Scanning;

namespace Compiler;

public class Compiler
{
    private ErrorHandler _errorHandler = new();
    
    public void Run(string src)
    {
        Scanner scanner = new(src, _errorHandler);
        List<Token> tokens = scanner.ScanTokens();
        foreach (Token token in tokens)
        {
            Console.WriteLine(token);
        }
    }
    
    static void Main(string[] args)
    {
        Compiler compiler = new();
        compiler.Run("var name = \"John\"");
    }
}