# $NetBSD: Makefile,v 1.59 2018/04/17 22:29:34 wiz Exp $
#

PKGREVISION=	15
COMMENT=	LibPrelude libraries and c++ easy bindings
LICENSE=	gnu-gpl-v2

.include "../../security/libprelude/Makefile.common"

CONF_FILES+=	${PREFIX}/share/examples/libprelude/client.conf ${PKG_SYSCONFDIR}/default/client.conf
CONF_FILES+=	${PREFIX}/share/examples/libprelude/global.conf ${PKG_SYSCONFDIR}/default/global.conf
CONF_FILES+=	${PREFIX}/share/examples/libprelude/idmef-client.conf ${PKG_SYSCONFDIR}/default/idmef-client.conf
CONF_FILES+=	${PREFIX}/share/examples/libprelude/tls.conf ${PKG_SYSCONFDIR}/default/tls.conf

CONFIGURE_ARGS+=	--without-python --without-lua --without-perl --without-ruby

# Work around mktime configure test failure leading to build issue
.if !empty(MACHINE_PLATFORM:MDarwin-*-x86_64)
CONFIGURE_ENV+=		ac_cv_func_working_mktime=yes
.endif

PKGCONFIG_OVERRIDE+=	src/libprelude.pc.in

PKG_OPTIONS_VAR=	PKG_OPTIONS.libprelude
PKG_SUPPORTED_OPTIONS=	debug

.include "../../mk/bsd.options.mk"

.if !empty(PKG_OPTIONS:Mdebug)
CFLAGS=			-O0 -ggdb
.endif

.include "../../devel/libltdl/convenience.mk"
.include "../../security/gnutls/buildlink3.mk"
.include "../../security/libgcrypt/buildlink3.mk"
.include "../../mk/pthread.buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
