# $NetBSD: Makefile,v 1.1 2018/06/12 23:50:02 minskim Exp $

DISTNAME=	glob-0.2.3
PKGNAME=	go-${DISTNAME}
MASTER_SITES=	${MASTER_SITE_GITHUB:=gobwas/}
CATEGORIES=	textproc
GITHUB_PROJECT=	${PKGBASE:S/^go-//}
GITHUB_TAG=	v${PKGVERSION_NOREV}

MAINTAINER=	minskim@NetBSD.org
HOMEPAGE=	https://github.com/gobwas/glob
COMMENT=	Go glob
LICENSE=	mit

GO_DIST_BASE=	${DISTNAME}
GO_SRCPATH=	${HOMEPAGE:S,https://,,}

.include "../../lang/go/go-package.mk"
.include "../../mk/bsd.pkg.mk"
