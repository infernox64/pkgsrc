# $NetBSD: Makefile,v 1.48 2018/04/04 10:23:17 leot Exp $

DISTNAME=		pkgsrc-guide-${PKGVERSION}
CATEGORIES=		# empty
MASTER_SITES=		${MASTER_SITE_LOCAL}
DISTFILES=		htdocs-share-20180404.tar.gz

MAINTAINER=		pkgsrc-users@NetBSD.org
HOMEPAGE=		http://www.NetBSD.org/docs/pkgsrc/
COMMENT=		The pkgsrc guide

# To regenerate pkgsrc/doc/pkgsrc.{html,txt} and
# htdocs/docs/pkgsrc, just run "make regen".

PKGVERSION!=		date '+%Y%m%d'
DIST_SUBDIR=		${PKGBASE}
USE_LANGUAGES=		# empty
MAKE_ENV+=		SED=${SED:Q}

PLIST_VARS=		ascii html pdf

DOCDIR=			${PREFIX}/share/doc/pkgsrc
.if exists(/usr/cvs/htdocs)
HTDOCSDIR?=		/usr/cvs/htdocs
.endif
HTDOCSDIR?=		${.CURDIR}/../../../htdocs

BUILD_DEFS+=		OUTPUTS
OUTPUTS?=		lint html html-split ascii pdf

INSTALLATION_DIRS=	${DOCDIR}

SUBST_CLASSES+=		docbook45
SUBST_STAGE.docbook45=	pre-configure
SUBST_FILES.docbook45=	${WRKDIR}/htdocs/share/xml/*
SUBST_SED.docbook45=	-e 's,V4\.2,V4.5,'
SUBST_SED.docbook45+=	-e 's,/4\.2/,/4.5/,'

.if defined(OUTPUTS)
.  if !empty(OUTPUTS:Mascii)
# the html is needed to build the ascii version.
OUTPUTS+=		html
PLIST.ascii=		yes
.  endif

.  if !empty(OUTPUTS:Mhtml) || !empty(OUTPUTS:Mhtml-split)
PLIST.html=		yes
.endif

.  if !empty(OUTPUTS:Mpdf)
PLIST.pdf=		yes
.endif

# only override the Makefile.common default if it is explicitly set
_GUIDE_OUTPUTS=		${OUTPUTS}
.endif

.include "Makefile.common"

# The source files are only symlinked into the WRKSRC, so that they can
# be easily modified, should the "lint" phase fail.
pre-extract:
	${MKDIR} ${WRKSRC}
	${LN} -s ${FILESDIR}/* ${WRKSRC}

do-build:
.for _output_ in ${OUTPUTS}
	@${ECHO} "-----> Building ${_output_} output"
	@cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${MAKE_PROGRAM} ${MAKE_FLAGS} ${_output_}
.endfor

post-build:
	@${RM} -f ${WRKSRC}/pkgsrc.tmp.html

do-install:
.if !empty(OUTPUTS:Mhtml) || !empty(OUTPUTS:Mhtml-split)
	${INSTALL_DATA} ${WRKSRC}/*.html ${DESTDIR}${DOCDIR}
	${INSTALL_DATA} ${WRKDIR}/htdocs/global.css ${DESTDIR}${DOCDIR}
.endif
.if !empty(OUTPUTS:Mascii)
	${INSTALL_DATA} ${WRKSRC}/pkgsrc.txt ${DESTDIR}${DOCDIR}
.endif
.if !empty(OUTPUTS:Mpdf)
	${INSTALL_DATA} ${WRKSRC}/pkgsrc.ps ${DESTDIR}${DOCDIR}
	${INSTALL_DATA} ${WRKSRC}/pkgsrc.pdf ${DESTDIR}${DOCDIR}
.endif

# install the single-file HTML and ascii output into the pkgsrc doc/
# directory, for distribution with pkgsrc.
#
# note that this uses ${CP} and not ${INSTALL_DATA} because the files
# stay in the development directories and have nothing to do with
# the permissions or ownership of installed files.
#
install-doc: build
	${CP} ${WRKSRC}/pkgsrc.html ${PKGSRCDIR}/doc
	${CP} ${WRKSRC}/pkgsrc.txt ${PKGSRCDIR}/doc

install-htdocs: build
	cd ${WRKSRC} && ${CP} pkgsrc.txt pkgsrc.pdf pkgsrc.ps *.html	\
		${HTDOCSDIR}/docs/pkgsrc

.PHONY: regen regen-doc regen-htdocs
regen: regen-doc regen-htdocs

regen-doc: build
	@${STEP_MSG} "Updating the files in pkgsrc/doc"
	cd .. && cvs update pkgsrc.*
	${MAKE} ${MAKEFLAGS} install-doc
	@${STEP_MSG} "Committing the files in pkgsrc/doc"
	cd .. && cvs commit -m "doc/pkgsrc.*: regen" pkgsrc.*

regen-htdocs: build
	@${STEP_MSG} "Updating the files in htdocs"
	cd ${HTDOCSDIR}/docs/pkgsrc && cvs update
	${MAKE} ${MAKEFLAGS} install-htdocs
	@${STEP_MSG} "Committing the files in htdocs"
	cd ${HTDOCSDIR}/docs/pkgsrc && cvs commit -m "regen"

do-lint:
	cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${MAKE_PROGRAM} lint

DETAB_PROGRAM= \
	if (/<programlisting>/ .. /<\/programlisting>/) {		\
		while (/(.*?)\t(.*)/s) {				\
			my $$filler = " " x (8 - (length($$1) % 8));	\
			$$_ = "$$1$$filler$$2";				\
		}							\
	}

# remove tabulators from the <programlisting> tags.
# TeX does not like them.
.PHONY: detab
detab:
	cd ${FILESDIR} &&						\
	perl -p -i".detab.tmp" -e '${DETAB_PROGRAM}' *.xml &&		\
	${RM} -f *.xml.detab.tmp

# Generates a new htdocs-share-*.tar.gz archive from the current
# contents of the htdocs directory. To make it available, it should
# be copied to ftp.NetBSD.org:/pub/pkgsrc/distfiles/LOCAL_PORTS/, i.e.:
#
# % scp htdocs-share-*.tar.gz ftp.NetBSD.org:/pub/pkgsrc/distfiles/LOCAL_PORTS/
# [... adjust DISTFILES variable ...]
# % make makesum
#
.PHONY: htdocs-share
htdocs-share:
	cd ${HTDOCSDIR}							\
	&& ${PAX} -wz -f ${.CURDIR}/htdocs-share-${PKGVERSION}.tar.gz	\
		-s ',^,htdocs/,'					\
		global.css share

.include "../../mk/bsd.pkg.mk"
