# $NetBSD: Makefile,v 1.93 2018/06/24 17:01:58 adam Exp $

DISTNAME=	w3c-libwww-5.4.2
PKGNAME=	${DISTNAME:S/^w3c-//}
CATEGORIES=	www devel
MASTER_SITES=	http://www.w3.org/Library/Distribution/
EXTRACT_SUFX=	.tgz

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://www.w3.org/Library/
COMMENT=	The W3C Reference Library
LICENSE=	w3c

USE_LIBTOOL=		yes
USE_TOOLS+=		gmake gzcat perl
GNU_CONFIGURE=		yes
CONFIGURE_ARGS+=	--with-expat=${BUILDLINK_PREFIX.expat}
CONFIGURE_ARGS+=	--with-ssl #=${SSLBASE:Q}
CONFIGURE_ARGS+=	--with-zlib
CONFIGURE_ARGS+=	--without-x

.include "../../mk/bsd.prefs.mk"

.if !empty(MACHINE_PLATFORM:MIRIX-5*)
# IRIX 5 does not have libregex
CONFIGURE_ARGS+=	--with-regex
.endif

PLIST_VARS+=		md5
.if ${OPSYS} == "SunOS"
# Solaris has a /usr/lib/libmd5 which seems to be enough, and
# /usr/local/lib/libmd5 breaks /usr/lib/librt
CONFIGURE_ARGS+=	--without-md5
.else
PLIST.md5=		yes
.  if !exists(/usr/include/md5.h)
CONFIGURE_ARGS+=	--with-md5=${SSLBASE:Q}
.  endif
.endif

.include "options.mk"

.include "../../devel/zlib/buildlink3.mk"
.include "../../security/openssl/buildlink3.mk"
.include "../../textproc/expat/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
