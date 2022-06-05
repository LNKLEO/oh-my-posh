module.exports = {
  docs: [
    {
      type: "category",
      label: "💡 Getting Started",
      collapsed: false,
      items: [
        "introduction",
        {
          type: "category",
          label: "🚀 Get started",
          collapsed: false,
          items: [
            {
              type: "category",
              label: "📦 Installation",
              collapsed: false,
              items: [
                "installation/windows",
                "installation/macos",
                "installation/linux",
              ],
            },
            "installation/fonts",
            "installation/prompt",
            "installation/customize",
          ],
        },
      ],
    },
    {
      type: "category",
      label: "⚙️ Configuration",
      items: [
        "configuration/overview",
        "configuration/block",
        "configuration/segment",
        "configuration/sample",
        "configuration/title",
        "configuration/colors",
        "configuration/templates",
        "configuration/secondary-prompt",
        "configuration/debug-prompt",
        "configuration/transient",
        "configuration/line-error",
        "configuration/tooltips",
      ],
    },
    {
      type: "category",
      label: "🌟 Segments",
      collapsed: true,
      items: [
        "segments/angular",
        "segments/aws",
        "segments/az",
        "segments/azfunc",
        "segments/battery",
        "segments/brewfather",
        "segments/cds",
        "segments/command",
        "segments/crystal",
        "segments/cf",
        "segments/cftarget",
        "segments/dart",
        "segments/dotnet",
        "segments/executiontime",
        "segments/exit",
        "segments/flutter",
        "segments/fossil",
        "segments/git",
        "segments/poshgit",
        "segments/golang",
        "segments/haskell",
        "segments/ipify",
        "segments/iterm",
        "segments/java",
        "segments/julia",
        "segments/kotlin",
        "segments/kubectl",
        "segments/nbgv",
        "segments/nightscout",
        "segments/npm",
        "segments/node",
        "segments/nx",
        "segments/os",
        "segments/owm",
        "segments/path",
        "segments/php",
        "segments/plastic",
        "segments/project",
        "segments/python",
        "segments/r",
        "segments/root",
        "segments/ruby",
        "segments/rust",
        "segments/session",
        "segments/shell",
        "segments/spotify",
        "segments/strava",
        "segments/svn",
        "segments/swift",
        "segments/sysinfo",
        "segments/terraform",
        "segments/text",
        "segments/time",
        "segments/ui5tooling",
        "segments/wakatime",
        "segments/wifi",
        "segments/winreg",
        "segments/ytm",
      ],
    },
    {
      type: "category",
      label: "🙋🏾‍♀️ Contributing",
      collapsed: true,
      items: [
        "contributing/started",
        "contributing/segment",
        "contributing/git",
        "contributing/plastic",
      ],
    },
    "themes",
    "share",
    "faq",
    "migrating",
    "contributors",
  ],
};
