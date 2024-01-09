using JetBrains.Annotations;
using Unii.Error;
using JetBrains.Annotations;

namespace Unii.Compile;

[PublicAPI]
public class Compiler
{
    private readonly ErrorHandler _errorHandler = new();
    
    // public void Run(string src)
    // {
    //     
    // }
}