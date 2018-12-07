using UnrealBuildTool;
using System.Collections.Generic;

public class MarsettlerEditorTarget : TargetRules
{
	public MarsettlerEditorTarget(TargetInfo Target) : base(Target)
	{
		Type = TargetType.Editor;

		ExtraModuleNames.AddRange( new string[] { "Marsettler" } );
	}
}
