# $NetBSD: Makefile,v 1.19 2015/07/27 22:22:18 wiz Exp $

DISTNAME=		xzip182
PKGNAME=		xzip-1.8.2
CATEGORIES=		games x11
MASTER_SITES=		http://www.eblong.com/zarf/ftp/
EXTRACT_SUFX=		.tar.Z

MAINTAINER=		pkgsrc-users@NetBSD.org
HOMEPAGE=		http://www.eblong.com/zarf/xzip.html
COMMENT=		Z-code interpreter for X11 (Infocom game format)

WRKSRC=			${WRKDIR}/xzip
BUILD_TARGET=		xzip
INSTALLATION_DIRS=	bin ${PKGMANDIR}/man6

.include "../../x11/libX11/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
