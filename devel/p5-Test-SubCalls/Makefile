# $NetBSD: Makefile,v 1.16 2017/12/30 14:16:28 wiz Exp $

DISTNAME=	Test-SubCalls-1.10
PKGNAME=	p5-${DISTNAME}
CATEGORIES=	devel perl5
MASTER_SITES=	${MASTER_SITE_PERL_CPAN:=Test/}

MAINTAINER=	pkgsrc-users@NetBSD.org
HOMEPAGE=	http://search.cpan.org/dist/Test-SubCalls/
COMMENT=	Perl5 module to track the number of times subs are called

DEPENDS+=	p5-Hook-LexWrap>=0.20:../../devel/p5-Hook-LexWrap

MAKE_ENV+=	PERL_USE_UNSAFE_INC=1
PERL5_PACKLIST=	auto/Test/SubCalls/.packlist

.include "../../lang/perl5/module.mk"
.include "../../mk/bsd.pkg.mk"
