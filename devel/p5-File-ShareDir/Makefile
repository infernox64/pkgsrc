# $NetBSD: Makefile,v 1.18 2018/06/24 06:58:15 wiz Exp $

DISTNAME=	File-ShareDir-1.114
PKGNAME=	p5-${DISTNAME}
CATEGORIES=	devel perl5
MASTER_SITES=	${MASTER_SITE_PERL_CPAN:=File/}

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://search.cpan.org/dist/File-ShareDir/
COMMENT=	Get a generic path to the share directory
LICENSE=	${PERL5_LICENSE}

DEPENDS+=	p5-Class-Inspector>=1.12:../../devel/p5-Class-Inspector

USE_LANGUAGES=		# empty
PERL5_MODULE_TYPE=	Module::Install::Bundled
PERL5_PACKLIST=		auto/File/ShareDir/.packlist

.include "../../lang/perl5/module.mk"
.include "../../mk/bsd.pkg.mk"
