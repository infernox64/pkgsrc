# $NetBSD: Makefile,v 1.17 2018/04/09 07:48:54 jaapb Exp $
#

VERSION=	0.8.1
GITHUB_PROJECT=	deriving
DISTNAME=	${GITHUB_PROJECT}-${VERSION}
PKGNAME=	ocaml-${GITHUB_PROJECT}-ocsigen-${VERSION}
COMMENT=	Extension to OCaml for deriving functions from type declarations
CATEGORIES=	devel
MASTER_SITES=	${MASTER_SITE_GITHUB:=ocsigen/}

MAINTAINER=	jaapb@NetBSD.org
HOMEPAGE=	https://github.com/ocsigen/deriving
LICENSE=	mit

USE_TOOLS=	gmake
OCAML_USE_OASIS_DYNRUN=	yes
OCAML_FINDLIB_DIRS=	deriving

.include "../../lang/camlp4/buildlink3.mk"
.include "../../math/ocaml-num/buildlink3.mk"
.include "../../mk/ocaml.mk"
.include "../../mk/bsd.pkg.mk"
