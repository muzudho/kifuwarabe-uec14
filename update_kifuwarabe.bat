@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo コピー中だぜ（＾～＾）...
copy "C:\Users\muzud\go\src\github.com\muzudho\kifuwarabe-uec14\kifuwarabe-uec14.exe" "C:\Users\muzud\OneDrive\ドキュメント\Tools\kifuwarabe-uec14\kifuwarabe-uec14.exe"
if %errorlevel%==0 (
    echo コピー完了！ よしよし（＾▽＾）
) else (
    echo エラー出たぜ... 確認してな。
)
pause
