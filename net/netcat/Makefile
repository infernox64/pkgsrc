# $NetBSD: Makefile,v 1.38 2014/10/09 14:06:45 wiz Exp $

DISTNAME=	nc110
PKGNAME=	netcat-1.10
PKGREVISION=	3
CATEGORIES=	net
MASTER_SITES=	ftp://coast.cs.purdue.edu/pub/tools/unix/netutils/netcat/
EXTRACT_SUFX=	.tgz

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://nc110.sourceforge.net/
COMMENT=	Read and write data across network connections

WRKSRC=		${WRKDIR}
MAKE_ENV+=	DFLAGS="-DGAPING_SECURITY_HOLE -DTELNET"

.include "../../mk/bsd.prefs.mk"

NC_TARGET.${OPSYS}?=	${LOWER_OPSYS}
NC_TARGET.IRIX=		irix
NC_TARGET.Interix=	netbsd
NC_TARGET.OpenBSD=	netbsd
NC_TARGET.DragonFly=	freebsd
NC_TARGET.Darwin=	generic
BUILD_TARGET=		${NC_TARGET.${OPSYS}}

EXTRA_PROGS=	data rservice xor
EXAMPLE_DIR=	share/examples/nc

PLIST_SRC=	${PKGDIR}/PLIST ${WRKSRC}/data/PLIST.data

INSTALLATION_DIRS=	${PKGMANDIR}/man1 sbin

post-build:
	@cd ${WRKSRC}/data && ${SETENV} ${MAKE_ENV} ${MAKE} ${EXTRA_PROGS}

# Note: the nc binary is installed into sbin to avoid a conflict with
# editors/nedit.
do-install:
	${INSTALL_PROGRAM} ${WRKSRC}/nc ${DESTDIR:Q}${PREFIX:Q}/sbin
.for i in ${EXTRA_PROGS}
	${INSTALL_PROGRAM} ${WRKSRC}/data/${i} \
		${DESTDIR:Q}${PREFIX:Q}/sbin/nc-${i}
.endfor
	${INSTALL_DATA_DIR} ${DESTDIR:Q}${PREFIX:Q}/${EXAMPLE_DIR}
	@(								\
	cd ${WRKSRC}/data;						\
	${RM} -f PLIST.data;						\
	for i in README Makefile *.c *.d;				\
	do								\
	    ${INSTALL_DATA} $$i ${DESTDIR:Q}${PREFIX:Q}/${EXAMPLE_DIR};	\
	    ${ECHO} ${EXAMPLE_DIR}/$$i >> PLIST.data;			\
	done;								\
	)
	${INSTALL_MAN} ${FILESDIR}/nc.1 \
		${DESTDIR:Q}${PREFIX:Q}/${PKGMANDIR}/man1/nc.1

.include "../../mk/bsd.pkg.mk"
