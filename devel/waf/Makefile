# $NetBSD: Makefile,v 1.39 2018/05/27 09:51:03 wiz Exp $

DISTNAME=	waf-2.0.8
CATEGORIES=	devel
MASTER_SITES=	https://waf.io/
EXTRACT_SUFX=	.tar.bz2

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	https://waf.io/
COMMENT=	Framework for configuring, compiling, and installing applications
LICENSE=	modified-bsd

REPLACE_PYTHON=		waf-light
INSTALLATION_DIRS=	bin

do-build:
	cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ./waf-light configure build

do-install:
	${INSTALL_SCRIPT} ${WRKSRC}/waf ${DESTDIR}${PREFIX}/bin

.include "../../lang/python/application.mk"
.include "../../mk/bsd.pkg.mk"
