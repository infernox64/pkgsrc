# $NetBSD: Makefile,v 1.27 2018/06/19 07:21:23 adam Exp $

DISTNAME=	qtbase-everywhere-src-${QTVERSION}
PKGNAME=	qt5-odbc-${QTVERSION}
COMMENT=	Qt5 odbc plugin

BUILD_DIRS=		src/plugins/sqldrivers
MAKE_ENV+=		QTDIR=${QTDIR}
INSTALL_ENV+=		INSTALL_ROOT=${DESTDIR}

LIBS+=			-lltdl

.include "../../x11/qt5-qtbase/Makefile.common"

do-configure:
	${RM} -f -r ${WRKSRC}/mkspecs
	cd ${WRKSRC}/${BUILD_DIRS} && \
	${QTDIR}/bin/qmake -o Makefile "QT.sql.enabled_features+=sql-odbc" "QMAKE_LIBS_ODBC=-lodbc"

post-install:
	rm -f ${DESTDIR}${PREFIX}/qt5/lib/cmake/Qt5Sql/Qt5Sql_QSQLiteDriverPlugin.cmake
	rm -f ${DESTDIR}${PREFIX}/qt5/plugins/sqldrivers/libqsqlite.*

.include "../../devel/libltdl/buildlink3.mk"
.include "../../databases/unixodbc/buildlink3.mk"
.include "../../x11/qt5-qtbase/buildlink3.mk"
.include "../../mk/bsd.pkg.mk"
