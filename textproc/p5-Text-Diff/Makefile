# $NetBSD: Makefile,v 1.27 2017/09/04 12:38:02 wiz Exp $

DISTNAME=	Text-Diff-1.45
PKGNAME=	p5-${DISTNAME}
CATEGORIES=	perl5 textproc
MASTER_SITES=	${MASTER_SITE_PERL_CPAN:=Text/}

MAINTAINER=	pkgsrc-users@NetBSD.org
COMMENT=	High-level text diffing module for Perl
HOMEPAGE=	http://search.cpan.org/dist/Text-Diff/
LICENSE=	${PERL5_LICENSE}

DEPENDS+=		p5-Algorithm-Diff>=1.19:../../devel/p5-Algorithm-Diff

USE_LANGUAGES=		# empty

PERL5_PACKLIST=		auto/Text/Diff/.packlist

.include "../../lang/perl5/module.mk"
.include "../../mk/bsd.pkg.mk"
