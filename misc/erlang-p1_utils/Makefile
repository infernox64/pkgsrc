# $NetBSD: Makefile,v 1.13 2018/03/24 14:24:00 fhajny Exp $

DISTNAME=	p1_utils-1.0.11
CATEGORIES=	misc
MASTER_SITES=	${MASTER_SITE_GITHUB:=processone/}

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	https://github.com/processone/p1_utils/
COMMENT=	Erlang Utility Modules from ProcessOne
LICENSE=	apache-2.0

GITHUB_PROJECT=	p1_utils

# Leave a non-versioned symlink in place
post-install:
	${LN} -sf ${DISTNAME} ${DESTDIR}${PREFIX}/lib/erlang/lib/p1_utils

.include "../../lang/erlang/module.mk"
.include "../../mk/bsd.pkg.mk"
