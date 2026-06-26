# 🐾 Virtual Pet Game (CLI)

A tiny, cozy virtual pet that lives in your terminal! Feed it, check on it, and watch it grow — all from the command line. 🌱

```
╔===================================================================╗
║       🐾  VIRTUAL PET GAME CLI  🐾                                ║
╚===================================================================╝
```

## ✨ What is this?

A little Tamagotchi-style pet simulator written in Go. Your pet's state (hunger, energy, life, name...) is saved to a local file, so it remembers you between sessions — just like a real little buddy waiting for you to come back. 💛

## 🚀 Getting Started

Make sure you have Go installed, then simply run:

```bash
go run .
```

That's it — your adventure begins!

## 🛠️ Available Options

| Flag | Description |
|------|-------------|
| `--name <string>` | 💡 Sets your pet's name (e.g. `--name=pamuk`) |
| `--status` | 📊 Shows your pet's current stats and exits |
| `--eat <1-8>` | 🍖 Feeds your pet a yummy treat from the market |
| `--savefile <path>` | 💾 Use a custom path for your save file |
| `--version` | 🔢 Prints the app's version |

### 🍽️ Menu (for `--eat`)

| # | Food |
|---|------|
| 1 | Omlet 🍳 |
| 2 | Fish 🐟 |
| 3 | Meat 🥩 |
| 4 | Apple 🍎 |
| 5 | Spaghetti 🍝 |
| 6 | Pizza 🍕 |
| 7 | Hamburger 🍔 |
| 8 | Chips 🍟 |

## 💬 Examples

```bash
go run ./src  --name=boncuk --status
go run ./src  --eat=1
```

## ⚠️ A Gentle Warning About Flag Order & Combinations

Here's the cozy little catch you need to know before you start playing: **the order and combination of your flags matters**, because of how the pet is set up the very first time. 🐣

When you run the app for the very first time (no save file exists yet), your pet has **no name**. The program checks for a name *before* it lets you do anything else — including feeding or checking status. So:

- ❌ `go run ./src  --eat=1` → on a brand-new pet, this will print `pls: enter name` and exit immediately. Your poor pet doesn't even get to eat its first meal!
- ❌ `go run ./src  --status` → same story, it'll exit before showing any stats, because there's no name yet.
- ✅ `go run ./src  --name=boncuk` → this works! Your pet hatches with a name. 🐾
- ✅ `go run ./src  --name=boncuk --eat=1` → this works too — name gets set *first*, then your pet happily eats. 🍳

**The rule of thumb:** always include `--name` together with your other flags until your pet has been named at least once. After that first successful run, the name is saved, so on every future run you're free to use `--eat`, `--status`, etc. on their own. 🎉

```bash
# First run ever — name your pet!
go run . --name=pamuk

# Every run after that — no need for --name anymore
go run ./src  --eat=3
go run ./src  --status
```

Think of it like adopting a pet in real life — you have to give it a name before you can feed it. 🐶💕

## 💾 Save File

By default, your pet's data is saved to a system-appropriate location. Want to keep multiple pets or store the save somewhere specific? Use:

```bash
go run ./src  --name=pamuk --savefile=./my-pet-save.json
```

## 🐛 Known Quirk

Because of the name-check happening early in the program, mixing `--eat`/`--status` without `--name` on a fresh save file will exit before doing anything. This isn't a bug in your setup — it's just how the current logic flows, so now you know the trick! 😉

## 📜 License

Made with 🧡 for anyone who just wants a tiny digital friend. (MIT)