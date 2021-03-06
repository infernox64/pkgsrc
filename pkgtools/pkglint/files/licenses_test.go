package main

import (
	"gopkg.in/check.v1"
)

func (s *Suite) Test_checklineLicense(c *check.C) {
	t := s.Init(c)

	t.SetupFileLines("licenses/gnu-gpl-v2",
		"Most software \u2026")
	mkline := t.NewMkLine("Makefile", 7, "LICENSE=dummy")
	G.CurrentDir = t.TmpDir()

	licenseChecker := &LicenseChecker{mkline}
	licenseChecker.Check("gpl-v2", opAssign)

	t.CheckOutputLines(
		"WARN: Makefile:7: License file ~/licenses/gpl-v2 does not exist.")

	licenseChecker.Check("no-profit shareware", opAssign)

	t.CheckOutputLines(
		"ERROR: Makefile:7: Parse error for license condition \"no-profit shareware\".")

	licenseChecker.Check("no-profit AND shareware", opAssign)

	t.CheckOutputLines(
		"WARN: Makefile:7: License file ~/licenses/no-profit does not exist.",
		"ERROR: Makefile:7: License \"no-profit\" must not be used.",
		"WARN: Makefile:7: License file ~/licenses/shareware does not exist.",
		"ERROR: Makefile:7: License \"shareware\" must not be used.")

	licenseChecker.Check("gnu-gpl-v2", opAssign)

	t.CheckOutputEmpty()

	licenseChecker.Check("gnu-gpl-v2 AND gnu-gpl-v2 OR gnu-gpl-v2", opAssign)

	t.CheckOutputLines(
		"ERROR: Makefile:7: AND and OR operators in license conditions can only be combined using parentheses.")

	licenseChecker.Check("(gnu-gpl-v2 OR gnu-gpl-v2) AND gnu-gpl-v2", opAssign)

	t.CheckOutputEmpty()
}

func (s *Suite) Test_checkToplevelUnusedLicenses(c *check.C) {
	t := s.Init(c)

	t.SetupFileLines("mk/bsd.pkg.mk", "# dummy")
	t.SetupFileLines("mk/fetch/sites.mk", "# dummy")
	t.SetupFileLines("mk/defaults/options.description", "option\tdescription")
	t.SetupFileLines("doc/TODO")
	t.SetupFileLines("mk/defaults/mk.conf")
	t.SetupFileLines("mk/tools/bsd.tools.mk",
		".include \"actual-tools.mk\"")
	t.SetupFileLines("mk/tools/actual-tools.mk")
	t.SetupFileLines("mk/tools/defaults.mk")
	t.SetupFileLines("mk/bsd.prefs.mk")
	t.SetupFileLines("mk/misc/category.mk")
	t.SetupFileLines("licenses/2-clause-bsd")
	t.SetupFileLines("licenses/gnu-gpl-v3")

	t.SetupFileLines("Makefile",
		MkRcsID,
		"SUBDIR+=\tcategory")

	t.SetupFileLines("category/Makefile",
		MkRcsID,
		"COMMENT=\tExample category",
		"",
		"SUBDIR+=\tpackage",
		"",
		".include \"../mk/misc/category.mk\"")

	t.SetupFileLines("category/package/Makefile",
		MkRcsID,
		"CATEGORIES=\tcategory",
		"",
		"COMMENT=Example package",
		"LICENSE=\t2-clause-bsd",
		"NO_CHECKSUM=\tyes")
	t.SetupFileLines("category/package/PLIST",
		PlistRcsID,
		"bin/program")

	G.Main("pkglint", "-r", "-Cglobal", t.TmpDir())

	t.CheckOutputLines(
		"WARN: ~/licenses/gnu-gpl-v3: This license seems to be unused.",
		"0 errors and 1 warning found.")
}
