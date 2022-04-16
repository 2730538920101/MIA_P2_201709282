// Generated from /home/ec2-user/MIA_P2/Cmd.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CmdLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TOK_MKDISK=1, TOK_RMDISK=2, TOK_FDISK=3, TOK_MOUNT=4, TOK_MKFS=5, TOK_LOGIN=6, 
		TOK_LOGOUT=7, TOK_MKGRP=8, TOK_RMGRP=9, TOK_MKUSR=10, TOK_RMUSR=11, TOK_MKFILE=12, 
		TOK_MKDIR=13, TOK_EXEC=14, TOK_PAUSE=15, TOK_REP=16, TOK_PATH=17, TOK_FIT=18, 
		TOK_SIZE=19, TOK_UNIT=20, TOK_NAME=21, TOK_TYPE=22, TOK_ID=23, TOK_USUARIO=24, 
		TOK_PASSWORD=25, TOK_PWD=26, TOK_CONT=27, TOK_GRP=28, TOK_RUTA=29, TOK_R=30, 
		TOK_P=31, TOK_FIRST=32, TOK_WORST=33, TOK_BEST=34, TOK_KB=35, TOK_MB=36, 
		TOK_BYTES=37, TOK_PRIMARIA=38, TOK_EXTENDIDA=39, TOK_LOGICA=40, TOK_FAST=41, 
		TOK_FULL=42, TOK_CADENA=43, TOK_NUMERO=44, TOK_IDENTIFICADOR=45, TOK_CAMINO=46, 
		TOK_PALABRA=47, TOK_IGUAL=48, COMENTARIOS=49, WHITESPACE=50;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"SLASH", "COMDOB", "HASHTAG", "DASH", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", 
			"TOK_MOUNT", "TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", 
			"TOK_MKUSR", "TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", 
			"TOK_REP", "TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", 
			"TOK_TYPE", "TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", 
			"TOK_GRP", "TOK_RUTA", "TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", 
			"TOK_KB", "TOK_MB", "TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", 
			"TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", 
			"TOK_CAMINO", "TOK_PALABRA", "TOK_IGUAL", "COMENTARIOS", "WHITESPACE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", 
			"TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", 
			"TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", "TOK_PATH", 
			"TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", "TOK_ID", 
			"TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", "TOK_RUTA", 
			"TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", "TOK_KB", "TOK_MB", 
			"TOK_BYTES", "TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FAST", 
			"TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", "TOK_CAMINO", 
			"TOK_PALABRA", "TOK_IGUAL", "COMENTARIOS", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public CmdLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Cmd.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0183\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b"+
		"\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3"+
		"\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20"+
		"\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24\3\25"+
		"\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34"+
		"\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36"+
		"\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3"+
		"!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3"+
		"&\3&\3\'\3\'\3\'\3(\3(\3)\3)\3*\3*\3+\3+\3,\3,\3-\3-\3.\3.\3.\3.\3.\3"+
		"/\3/\3/\3/\3/\3\60\3\60\7\60\u014f\n\60\f\60\16\60\u0152\13\60\3\60\3"+
		"\60\3\61\6\61\u0157\n\61\r\61\16\61\u0158\3\62\3\62\3\62\3\62\3\62\3\63"+
		"\3\63\6\63\u0162\n\63\r\63\16\63\u0163\6\63\u0166\n\63\r\63\16\63\u0167"+
		"\3\64\6\64\u016b\n\64\r\64\16\64\u016c\3\65\3\65\3\66\3\66\7\66\u0173"+
		"\n\66\f\66\16\66\u0176\13\66\3\66\5\66\u0179\n\66\3\66\3\66\3\67\6\67"+
		"\u017e\n\67\r\67\16\67\u017f\3\67\3\67\2\28\3\2\5\2\7\2\t\2\13\3\r\4\17"+
		"\5\21\6\23\7\25\b\27\t\31\n\33\13\35\f\37\r!\16#\17%\20\'\21)\22+\23-"+
		"\24/\25\61\26\63\27\65\30\67\319\32;\33=\34?\35A\36C\37E G!I\"K#M$O%Q"+
		"&S\'U(W)Y*[+],_-a.c/e\60g\61i\62k\63m\64\3\2!\4\2OOoo\4\2MMmm\4\2FFff"+
		"\4\2KKkk\4\2UUuu\4\2TTtt\4\2HHhh\4\2QQqq\4\2WWww\4\2PPpp\4\2VVvv\4\2N"+
		"Nnn\4\2IIii\4\2RRrr\4\2GGgg\4\2ZZzz\4\2EEee\4\2CCcc\4\2JJjj\4\2\\\\||"+
		"\4\2[[{{\4\2YYyy\4\2DDdd\4\2\f\f$$\3\2\62;\5\2CC\\\\c|\5\2\f\f\"\"$$\7"+
		"\2/\60\62;C\\aac|\3\2??\3\2\f\f\5\2\13\f\17\17\"\"\2\u0186\2\13\3\2\2"+
		"\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2"+
		"\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2"+
		"\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2"+
		"\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S"+
		"\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2"+
		"\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2"+
		"\2m\3\2\2\2\3o\3\2\2\2\5q\3\2\2\2\7s\3\2\2\2\tu\3\2\2\2\13w\3\2\2\2\r"+
		"~\3\2\2\2\17\u0085\3\2\2\2\21\u008b\3\2\2\2\23\u0091\3\2\2\2\25\u0096"+
		"\3\2\2\2\27\u009c\3\2\2\2\31\u00a3\3\2\2\2\33\u00a9\3\2\2\2\35\u00af\3"+
		"\2\2\2\37\u00b5\3\2\2\2!\u00bb\3\2\2\2#\u00c2\3\2\2\2%\u00c8\3\2\2\2\'"+
		"\u00cd\3\2\2\2)\u00d3\3\2\2\2+\u00d7\3\2\2\2-\u00dd\3\2\2\2/\u00e2\3\2"+
		"\2\2\61\u00e8\3\2\2\2\63\u00ee\3\2\2\2\65\u00f4\3\2\2\2\67\u00fa\3\2\2"+
		"\29\u00fe\3\2\2\2;\u0107\3\2\2\2=\u0111\3\2\2\2?\u0116\3\2\2\2A\u011c"+
		"\3\2\2\2C\u0121\3\2\2\2E\u0127\3\2\2\2G\u012a\3\2\2\2I\u012d\3\2\2\2K"+
		"\u0130\3\2\2\2M\u0133\3\2\2\2O\u0136\3\2\2\2Q\u0138\3\2\2\2S\u013a\3\2"+
		"\2\2U\u013c\3\2\2\2W\u013e\3\2\2\2Y\u0140\3\2\2\2[\u0142\3\2\2\2]\u0147"+
		"\3\2\2\2_\u014c\3\2\2\2a\u0156\3\2\2\2c\u015a\3\2\2\2e\u0165\3\2\2\2g"+
		"\u016a\3\2\2\2i\u016e\3\2\2\2k\u0170\3\2\2\2m\u017d\3\2\2\2op\7\61\2\2"+
		"p\4\3\2\2\2qr\7$\2\2r\6\3\2\2\2st\7%\2\2t\b\3\2\2\2uv\7/\2\2v\n\3\2\2"+
		"\2wx\t\2\2\2xy\t\3\2\2yz\t\4\2\2z{\t\5\2\2{|\t\6\2\2|}\t\3\2\2}\f\3\2"+
		"\2\2~\177\t\7\2\2\177\u0080\t\2\2\2\u0080\u0081\t\4\2\2\u0081\u0082\t"+
		"\5\2\2\u0082\u0083\t\6\2\2\u0083\u0084\t\3\2\2\u0084\16\3\2\2\2\u0085"+
		"\u0086\t\b\2\2\u0086\u0087\t\4\2\2\u0087\u0088\t\5\2\2\u0088\u0089\t\6"+
		"\2\2\u0089\u008a\t\3\2\2\u008a\20\3\2\2\2\u008b\u008c\t\2\2\2\u008c\u008d"+
		"\t\t\2\2\u008d\u008e\t\n\2\2\u008e\u008f\t\13\2\2\u008f\u0090\t\f\2\2"+
		"\u0090\22\3\2\2\2\u0091\u0092\t\2\2\2\u0092\u0093\t\3\2\2\u0093\u0094"+
		"\t\b\2\2\u0094\u0095\t\6\2\2\u0095\24\3\2\2\2\u0096\u0097\t\r\2\2\u0097"+
		"\u0098\t\t\2\2\u0098\u0099\t\16\2\2\u0099\u009a\t\5\2\2\u009a\u009b\t"+
		"\13\2\2\u009b\26\3\2\2\2\u009c\u009d\t\r\2\2\u009d\u009e\t\t\2\2\u009e"+
		"\u009f\t\16\2\2\u009f\u00a0\t\t\2\2\u00a0\u00a1\t\n\2\2\u00a1\u00a2\t"+
		"\f\2\2\u00a2\30\3\2\2\2\u00a3\u00a4\t\2\2\2\u00a4\u00a5\t\3\2\2\u00a5"+
		"\u00a6\t\16\2\2\u00a6\u00a7\t\7\2\2\u00a7\u00a8\t\17\2\2\u00a8\32\3\2"+
		"\2\2\u00a9\u00aa\t\7\2\2\u00aa\u00ab\t\2\2\2\u00ab\u00ac\t\16\2\2\u00ac"+
		"\u00ad\t\7\2\2\u00ad\u00ae\t\17\2\2\u00ae\34\3\2\2\2\u00af\u00b0\t\2\2"+
		"\2\u00b0\u00b1\t\3\2\2\u00b1\u00b2\t\n\2\2\u00b2\u00b3\t\6\2\2\u00b3\u00b4"+
		"\t\7\2\2\u00b4\36\3\2\2\2\u00b5\u00b6\t\7\2\2\u00b6\u00b7\t\2\2\2\u00b7"+
		"\u00b8\t\n\2\2\u00b8\u00b9\t\6\2\2\u00b9\u00ba\t\7\2\2\u00ba \3\2\2\2"+
		"\u00bb\u00bc\t\2\2\2\u00bc\u00bd\t\3\2\2\u00bd\u00be\t\b\2\2\u00be\u00bf"+
		"\t\5\2\2\u00bf\u00c0\t\r\2\2\u00c0\u00c1\t\20\2\2\u00c1\"\3\2\2\2\u00c2"+
		"\u00c3\t\2\2\2\u00c3\u00c4\t\3\2\2\u00c4\u00c5\t\4\2\2\u00c5\u00c6\t\5"+
		"\2\2\u00c6\u00c7\t\7\2\2\u00c7$\3\2\2\2\u00c8\u00c9\t\20\2\2\u00c9\u00ca"+
		"\t\21\2\2\u00ca\u00cb\t\20\2\2\u00cb\u00cc\t\22\2\2\u00cc&\3\2\2\2\u00cd"+
		"\u00ce\t\17\2\2\u00ce\u00cf\t\23\2\2\u00cf\u00d0\t\n\2\2\u00d0\u00d1\t"+
		"\6\2\2\u00d1\u00d2\t\20\2\2\u00d2(\3\2\2\2\u00d3\u00d4\t\7\2\2\u00d4\u00d5"+
		"\t\20\2\2\u00d5\u00d6\t\17\2\2\u00d6*\3\2\2\2\u00d7\u00d8\5\t\5\2\u00d8"+
		"\u00d9\t\17\2\2\u00d9\u00da\t\23\2\2\u00da\u00db\t\f\2\2\u00db\u00dc\t"+
		"\24\2\2\u00dc,\3\2\2\2\u00dd\u00de\5\t\5\2\u00de\u00df\t\b\2\2\u00df\u00e0"+
		"\t\5\2\2\u00e0\u00e1\t\f\2\2\u00e1.\3\2\2\2\u00e2\u00e3\5\t\5\2\u00e3"+
		"\u00e4\t\6\2\2\u00e4\u00e5\t\5\2\2\u00e5\u00e6\t\25\2\2\u00e6\u00e7\t"+
		"\20\2\2\u00e7\60\3\2\2\2\u00e8\u00e9\5\t\5\2\u00e9\u00ea\t\n\2\2\u00ea"+
		"\u00eb\t\13\2\2\u00eb\u00ec\t\5\2\2\u00ec\u00ed\t\f\2\2\u00ed\62\3\2\2"+
		"\2\u00ee\u00ef\5\t\5\2\u00ef\u00f0\t\13\2\2\u00f0\u00f1\t\23\2\2\u00f1"+
		"\u00f2\t\2\2\2\u00f2\u00f3\t\20\2\2\u00f3\64\3\2\2\2\u00f4\u00f5\5\t\5"+
		"\2\u00f5\u00f6\t\f\2\2\u00f6\u00f7\t\26\2\2\u00f7\u00f8\t\17\2\2\u00f8"+
		"\u00f9\t\20\2\2\u00f9\66\3\2\2\2\u00fa\u00fb\5\t\5\2\u00fb\u00fc\t\5\2"+
		"\2\u00fc\u00fd\t\4\2\2\u00fd8\3\2\2\2\u00fe\u00ff\5\t\5\2\u00ff\u0100"+
		"\t\n\2\2\u0100\u0101\t\6\2\2\u0101\u0102\t\n\2\2\u0102\u0103\t\23\2\2"+
		"\u0103\u0104\t\7\2\2\u0104\u0105\t\5\2\2\u0105\u0106\t\t\2\2\u0106:\3"+
		"\2\2\2\u0107\u0108\5\t\5\2\u0108\u0109\t\17\2\2\u0109\u010a\t\23\2\2\u010a"+
		"\u010b\t\6\2\2\u010b\u010c\t\6\2\2\u010c\u010d\t\27\2\2\u010d\u010e\t"+
		"\t\2\2\u010e\u010f\t\7\2\2\u010f\u0110\t\4\2\2\u0110<\3\2\2\2\u0111\u0112"+
		"\5\t\5\2\u0112\u0113\t\17\2\2\u0113\u0114\t\27\2\2\u0114\u0115\t\4\2\2"+
		"\u0115>\3\2\2\2\u0116\u0117\5\t\5\2\u0117\u0118\t\22\2\2\u0118\u0119\t"+
		"\t\2\2\u0119\u011a\t\13\2\2\u011a\u011b\t\f\2\2\u011b@\3\2\2\2\u011c\u011d"+
		"\5\t\5\2\u011d\u011e\t\16\2\2\u011e\u011f\t\7\2\2\u011f\u0120\t\17\2\2"+
		"\u0120B\3\2\2\2\u0121\u0122\5\t\5\2\u0122\u0123\t\7\2\2\u0123\u0124\t"+
		"\n\2\2\u0124\u0125\t\f\2\2\u0125\u0126\t\23\2\2\u0126D\3\2\2\2\u0127\u0128"+
		"\5\t\5\2\u0128\u0129\t\7\2\2\u0129F\3\2\2\2\u012a\u012b\5\t\5\2\u012b"+
		"\u012c\t\17\2\2\u012cH\3\2\2\2\u012d\u012e\t\b\2\2\u012e\u012f\t\b\2\2"+
		"\u012fJ\3\2\2\2\u0130\u0131\t\27\2\2\u0131\u0132\t\b\2\2\u0132L\3\2\2"+
		"\2\u0133\u0134\t\30\2\2\u0134\u0135\t\b\2\2\u0135N\3\2\2\2\u0136\u0137"+
		"\t\3\2\2\u0137P\3\2\2\2\u0138\u0139\t\2\2\2\u0139R\3\2\2\2\u013a\u013b"+
		"\t\30\2\2\u013bT\3\2\2\2\u013c\u013d\t\17\2\2\u013dV\3\2\2\2\u013e\u013f"+
		"\t\20\2\2\u013fX\3\2\2\2\u0140\u0141\t\r\2\2\u0141Z\3\2\2\2\u0142\u0143"+
		"\t\b\2\2\u0143\u0144\t\23\2\2\u0144\u0145\t\6\2\2\u0145\u0146\t\f\2\2"+
		"\u0146\\\3\2\2\2\u0147\u0148\t\b\2\2\u0148\u0149\t\n\2\2\u0149\u014a\t"+
		"\r\2\2\u014a\u014b\t\r\2\2\u014b^\3\2\2\2\u014c\u0150\5\5\3\2\u014d\u014f"+
		"\n\31\2\2\u014e\u014d\3\2\2\2\u014f\u0152\3\2\2\2\u0150\u014e\3\2\2\2"+
		"\u0150\u0151\3\2\2\2\u0151\u0153\3\2\2\2\u0152\u0150\3\2\2\2\u0153\u0154"+
		"\5\5\3\2\u0154`\3\2\2\2\u0155\u0157\t\32\2\2\u0156\u0155\3\2\2\2\u0157"+
		"\u0158\3\2\2\2\u0158\u0156\3\2\2\2\u0158\u0159\3\2\2\2\u0159b\3\2\2\2"+
		"\u015a\u015b\t\32\2\2\u015b\u015c\t\32\2\2\u015c\u015d\t\32\2\2\u015d"+
		"\u015e\t\33\2\2\u015ed\3\2\2\2\u015f\u0161\5\3\2\2\u0160\u0162\n\34\2"+
		"\2\u0161\u0160\3\2\2\2\u0162\u0163\3\2\2\2\u0163\u0161\3\2\2\2\u0163\u0164"+
		"\3\2\2\2\u0164\u0166\3\2\2\2\u0165\u015f\3\2\2\2\u0166\u0167\3\2\2\2\u0167"+
		"\u0165\3\2\2\2\u0167\u0168\3\2\2\2\u0168f\3\2\2\2\u0169\u016b\t\35\2\2"+
		"\u016a\u0169\3\2\2\2\u016b\u016c\3\2\2\2\u016c\u016a\3\2\2\2\u016c\u016d"+
		"\3\2\2\2\u016dh\3\2\2\2\u016e\u016f\t\36\2\2\u016fj\3\2\2\2\u0170\u0174"+
		"\5\7\4\2\u0171\u0173\n\37\2\2\u0172\u0171\3\2\2\2\u0173\u0176\3\2\2\2"+
		"\u0174\u0172\3\2\2\2\u0174\u0175\3\2\2\2\u0175\u0178\3\2\2\2\u0176\u0174"+
		"\3\2\2\2\u0177\u0179\t\37\2\2\u0178\u0177\3\2\2\2\u0178\u0179\3\2\2\2"+
		"\u0179\u017a\3\2\2\2\u017a\u017b\b\66\2\2\u017bl\3\2\2\2\u017c\u017e\t"+
		" \2\2\u017d\u017c\3\2\2\2\u017e\u017f\3\2\2\2\u017f\u017d\3\2\2\2\u017f"+
		"\u0180\3\2\2\2\u0180\u0181\3\2\2\2\u0181\u0182\b\67\2\2\u0182n\3\2\2\2"+
		"\13\2\u0150\u0158\u0163\u0167\u016c\u0174\u0178\u017f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}