# $NetBSD: Makefile,v 1.27 2018/04/17 22:29:36 wiz Exp $

MODNAME=		pdo_dblib
CATEGORIES+=		databases
COMMENT=		PHP extension for PHP Data Objects (FreeTDS/Sybase/MS SQL)
PKGREVISION=		1

USE_PHP_EXT_PATCHES=	yes

.include "../../mk/bsd.prefs.mk"

.include "../../lang/php/ext.mk"

CONFIGURE_ARGS+=	--with-${MODNAME}
CONFIGURE_ARGS+=	--with-${PKGMODNAME}=shared,${BUILDLINK_DIR}

# requires 0.63 version to compile, doesn't compile with 0.61
BUILDLINK_API_DEPENDS.freetds+=	freetds>=0.63

.include "../../databases/freetds/buildlink3.mk"
.include "../../databases/php-pdo/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
