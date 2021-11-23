module.exports = {
  docs: [
    {
      type: "category",
      label: "💡 Getting Started",
      collapsed: false,
      items: [
        "introduction",
        "upgrading",
        {
          type: "category",
          label: "🚀 Installation",
          collapsed: false,
          items: ["pwsh", "windows", "macos", "linux"],
        },
      ],
    },
    {
      type: "category",
      label: "⚙️ Configuration",
      items: [
        "config-overview",
        "config-block",
        "config-segment",
        "config-sample",
        "config-title",
        "config-colors",
        "config-text-style",
        "config-transient",
        "config-tooltips",
        "config-fonts"
      ],
    },
    {
      type: "category",
      label: "🌟 Segments",
      collapsed: true,
      items: [
        "angular",
        "aws",
        "az",
        "azfunc",
        "battery",
        "command",
        "crystal",
        "dart",
        "dotnet",
        "environment",
        "executiontime",
        "exit",
        "git",
        "poshgit",
        "golang",
        "java",
        "julia",
        "kubectl",
        "nbgv",
        "nightscout",
        "node",
        "os",
        "owm",
        "path",
        "php",
        "python",
        "root",
        "ruby",
        "rust",
        "session",
        "shell",
        "spotify",
        "sysinfo",
        "terraform",
        "text",
        "time",
        "ytm",
      ],
    },
    {
      type: "category",
      label: "🙋🏾‍♀️ Contributing",
      collapsed: true,
      items: [
        "contributing_started",
        "contributing_segment",
        "contributing_git",
      ],
    },
    "themes",
    "share",
    "faq",
    "contributors",
  ],
};
