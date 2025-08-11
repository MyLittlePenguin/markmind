{
  pkgs ? import <nixpkgs> { },
}:
pkgs.mkShell {
  name = "Good old Agis";
  buildInputs = [
    pkgs.tailwindcss_4
    pkgs.templ
  ];
  shellHook = ''
    echo "Start env for Good old Agis"
  '';
}
