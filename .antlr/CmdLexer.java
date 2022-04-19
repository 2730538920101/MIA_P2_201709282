// Generated from /home/ec2-user/MIA_P2/Cmd.g4 by ANTLR 4.8

        import "./parser"
        import "./util"
        import "./ast"
        import arrayList "github.com/colegno/arraylist"

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
		TOK_P=31, TOK_FIRST=32, TOK_WORST=33, TOK_BEST=34, TOK_FAST=35, TOK_FULL=36, 
		TOK_CADENA=37, TOK_NUMERO=38, TOK_IDENTIFICADOR=39, TOK_CAMINO=40, TOK_PALABRA=41, 
		COMENTARIOS=42, WHITESPACE=43;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"SLASH", "COMDOB", "HASHTAG", "DASH", "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", 
			"TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", 
			"TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", 
			"TOK_PAUSE", "TOK_REP", "TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", 
			"TOK_NAME", "TOK_TYPE", "TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", 
			"TOK_CONT", "TOK_GRP", "TOK_RUTA", "TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", 
			"TOK_BEST", "TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", 
			"TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE"
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
			"TOK_R", "TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", "TOK_FAST", "TOK_FULL", 
			"TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", "TOK_CAMINO", "TOK_PALABRA", 
			"COMENTARIOS", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2-\u0170\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\3\2\3\2\3\3\3\3\3\4\3\4\3\5"+
		"\3\5\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3"+
		"\20\3\21\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3"+
		"\23\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3"+
		"\25\3\25\3\25\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3"+
		"\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3"+
		"\32\3\32\3\33\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3"+
		"\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3"+
		"\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!\3!\3"+
		"!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3%\3%\3%\3&\3&"+
		"\3&\3\'\3\'\3\'\3(\3(\3(\3)\3)\3)\3)\3)\3*\3*\3*\3*\3*\3+\3+\7+\u0139"+
		"\n+\f+\16+\u013c\13+\3+\3+\3,\6,\u0141\n,\r,\16,\u0142\3-\3-\3-\3-\5-"+
		"\u0149\n-\3.\3.\7.\u014d\n.\f.\16.\u0150\13.\6.\u0152\n.\r.\16.\u0153"+
		"\3/\6/\u0157\n/\r/\16/\u0158\3\60\3\60\7\60\u015d\n\60\f\60\16\60\u0160"+
		"\13\60\3\60\7\60\u0163\n\60\f\60\16\60\u0166\13\60\3\60\3\60\3\61\6\61"+
		"\u016b\n\61\r\61\16\61\u016c\3\61\3\61\2\2\62\3\2\5\2\7\2\t\2\13\2\r\3"+
		"\17\4\21\5\23\6\25\7\27\b\31\t\33\n\35\13\37\f!\r#\16%\17\'\20)\21+\22"+
		"-\23/\24\61\25\63\26\65\27\67\309\31;\32=\33?\34A\35C\36E\37G I!K\"M#"+
		"O$Q%S&U\'W(Y)[*]+_,a-\3\2 \4\2OOoo\4\2MMmm\4\2FFff\4\2KKkk\4\2UUuu\4\2"+
		"TTtt\4\2HHhh\4\2QQqq\4\2WWww\4\2PPpp\4\2VVvv\4\2NNnn\4\2IIii\4\2RRrr\4"+
		"\2GGgg\4\2ZZzz\4\2EEee\4\2CCcc\4\2JJjj\4\2\\\\||\4\2[[{{\4\2YYyy\4\2D"+
		"Ddd\4\2\f\f$$\5\2\60\60\62\62;;\7\2\60\60CC\\\\cc||\5\2\f\f\"\"$$\n\2"+
		"/\60\62\62;;CC\\\\aacc||\3\2\f\f\5\2\13\f\17\17\"\"\2\u0172\2\r\3\2\2"+
		"\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2"+
		"\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2"+
		"\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2"+
		"\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2"+
		"\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2"+
		"\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U"+
		"\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2"+
		"\2\2\3c\3\2\2\2\5e\3\2\2\2\7g\3\2\2\2\ti\3\2\2\2\13k\3\2\2\2\rm\3\2\2"+
		"\2\17t\3\2\2\2\21{\3\2\2\2\23\u0081\3\2\2\2\25\u0087\3\2\2\2\27\u008c"+
		"\3\2\2\2\31\u0092\3\2\2\2\33\u0099\3\2\2\2\35\u009f\3\2\2\2\37\u00a5\3"+
		"\2\2\2!\u00ab\3\2\2\2#\u00b1\3\2\2\2%\u00b8\3\2\2\2\'\u00be\3\2\2\2)\u00c3"+
		"\3\2\2\2+\u00c9\3\2\2\2-\u00cd\3\2\2\2/\u00d3\3\2\2\2\61\u00d8\3\2\2\2"+
		"\63\u00de\3\2\2\2\65\u00e4\3\2\2\2\67\u00ea\3\2\2\29\u00f0\3\2\2\2;\u00f4"+
		"\3\2\2\2=\u00fd\3\2\2\2?\u0107\3\2\2\2A\u010c\3\2\2\2C\u0112\3\2\2\2E"+
		"\u0117\3\2\2\2G\u011d\3\2\2\2I\u0120\3\2\2\2K\u0123\3\2\2\2M\u0126\3\2"+
		"\2\2O\u0129\3\2\2\2Q\u012c\3\2\2\2S\u0131\3\2\2\2U\u0136\3\2\2\2W\u0140"+
		"\3\2\2\2Y\u0144\3\2\2\2[\u0151\3\2\2\2]\u0156\3\2\2\2_\u015a\3\2\2\2a"+
		"\u016a\3\2\2\2cd\7\61\2\2d\4\3\2\2\2ef\7$\2\2f\6\3\2\2\2gh\7%\2\2h\b\3"+
		"\2\2\2ij\7/\2\2j\n\3\2\2\2kl\7?\2\2l\f\3\2\2\2mn\t\2\2\2no\t\3\2\2op\t"+
		"\4\2\2pq\t\5\2\2qr\t\6\2\2rs\t\3\2\2s\16\3\2\2\2tu\t\7\2\2uv\t\2\2\2v"+
		"w\t\4\2\2wx\t\5\2\2xy\t\6\2\2yz\t\3\2\2z\20\3\2\2\2{|\t\b\2\2|}\t\4\2"+
		"\2}~\t\5\2\2~\177\t\6\2\2\177\u0080\t\3\2\2\u0080\22\3\2\2\2\u0081\u0082"+
		"\t\2\2\2\u0082\u0083\t\t\2\2\u0083\u0084\t\n\2\2\u0084\u0085\t\13\2\2"+
		"\u0085\u0086\t\f\2\2\u0086\24\3\2\2\2\u0087\u0088\t\2\2\2\u0088\u0089"+
		"\t\3\2\2\u0089\u008a\t\b\2\2\u008a\u008b\t\6\2\2\u008b\26\3\2\2\2\u008c"+
		"\u008d\t\r\2\2\u008d\u008e\t\t\2\2\u008e\u008f\t\16\2\2\u008f\u0090\t"+
		"\5\2\2\u0090\u0091\t\13\2\2\u0091\30\3\2\2\2\u0092\u0093\t\r\2\2\u0093"+
		"\u0094\t\t\2\2\u0094\u0095\t\16\2\2\u0095\u0096\t\t\2\2\u0096\u0097\t"+
		"\n\2\2\u0097\u0098\t\f\2\2\u0098\32\3\2\2\2\u0099\u009a\t\2\2\2\u009a"+
		"\u009b\t\3\2\2\u009b\u009c\t\16\2\2\u009c\u009d\t\7\2\2\u009d\u009e\t"+
		"\17\2\2\u009e\34\3\2\2\2\u009f\u00a0\t\7\2\2\u00a0\u00a1\t\2\2\2\u00a1"+
		"\u00a2\t\16\2\2\u00a2\u00a3\t\7\2\2\u00a3\u00a4\t\17\2\2\u00a4\36\3\2"+
		"\2\2\u00a5\u00a6\t\2\2\2\u00a6\u00a7\t\3\2\2\u00a7\u00a8\t\n\2\2\u00a8"+
		"\u00a9\t\6\2\2\u00a9\u00aa\t\7\2\2\u00aa \3\2\2\2\u00ab\u00ac\t\7\2\2"+
		"\u00ac\u00ad\t\2\2\2\u00ad\u00ae\t\n\2\2\u00ae\u00af\t\6\2\2\u00af\u00b0"+
		"\t\7\2\2\u00b0\"\3\2\2\2\u00b1\u00b2\t\2\2\2\u00b2\u00b3\t\3\2\2\u00b3"+
		"\u00b4\t\b\2\2\u00b4\u00b5\t\5\2\2\u00b5\u00b6\t\r\2\2\u00b6\u00b7\t\20"+
		"\2\2\u00b7$\3\2\2\2\u00b8\u00b9\t\2\2\2\u00b9\u00ba\t\3\2\2\u00ba\u00bb"+
		"\t\4\2\2\u00bb\u00bc\t\5\2\2\u00bc\u00bd\t\7\2\2\u00bd&\3\2\2\2\u00be"+
		"\u00bf\t\20\2\2\u00bf\u00c0\t\21\2\2\u00c0\u00c1\t\20\2\2\u00c1\u00c2"+
		"\t\22\2\2\u00c2(\3\2\2\2\u00c3\u00c4\t\17\2\2\u00c4\u00c5\t\23\2\2\u00c5"+
		"\u00c6\t\n\2\2\u00c6\u00c7\t\6\2\2\u00c7\u00c8\t\20\2\2\u00c8*\3\2\2\2"+
		"\u00c9\u00ca\t\7\2\2\u00ca\u00cb\t\20\2\2\u00cb\u00cc\t\17\2\2\u00cc,"+
		"\3\2\2\2\u00cd\u00ce\5\t\5\2\u00ce\u00cf\t\17\2\2\u00cf\u00d0\t\23\2\2"+
		"\u00d0\u00d1\t\f\2\2\u00d1\u00d2\t\24\2\2\u00d2.\3\2\2\2\u00d3\u00d4\5"+
		"\t\5\2\u00d4\u00d5\t\b\2\2\u00d5\u00d6\t\5\2\2\u00d6\u00d7\t\f\2\2\u00d7"+
		"\60\3\2\2\2\u00d8\u00d9\5\t\5\2\u00d9\u00da\t\6\2\2\u00da\u00db\t\5\2"+
		"\2\u00db\u00dc\t\25\2\2\u00dc\u00dd\t\20\2\2\u00dd\62\3\2\2\2\u00de\u00df"+
		"\5\t\5\2\u00df\u00e0\t\n\2\2\u00e0\u00e1\t\13\2\2\u00e1\u00e2\t\5\2\2"+
		"\u00e2\u00e3\t\f\2\2\u00e3\64\3\2\2\2\u00e4\u00e5\5\t\5\2\u00e5\u00e6"+
		"\t\13\2\2\u00e6\u00e7\t\23\2\2\u00e7\u00e8\t\2\2\2\u00e8\u00e9\t\20\2"+
		"\2\u00e9\66\3\2\2\2\u00ea\u00eb\5\t\5\2\u00eb\u00ec\t\f\2\2\u00ec\u00ed"+
		"\t\26\2\2\u00ed\u00ee\t\17\2\2\u00ee\u00ef\t\20\2\2\u00ef8\3\2\2\2\u00f0"+
		"\u00f1\5\t\5\2\u00f1\u00f2\t\5\2\2\u00f2\u00f3\t\4\2\2\u00f3:\3\2\2\2"+
		"\u00f4\u00f5\5\t\5\2\u00f5\u00f6\t\n\2\2\u00f6\u00f7\t\6\2\2\u00f7\u00f8"+
		"\t\n\2\2\u00f8\u00f9\t\23\2\2\u00f9\u00fa\t\7\2\2\u00fa\u00fb\t\5\2\2"+
		"\u00fb\u00fc\t\t\2\2\u00fc<\3\2\2\2\u00fd\u00fe\5\t\5\2\u00fe\u00ff\t"+
		"\17\2\2\u00ff\u0100\t\23\2\2\u0100\u0101\t\6\2\2\u0101\u0102\t\6\2\2\u0102"+
		"\u0103\t\27\2\2\u0103\u0104\t\t\2\2\u0104\u0105\t\7\2\2\u0105\u0106\t"+
		"\4\2\2\u0106>\3\2\2\2\u0107\u0108\5\t\5\2\u0108\u0109\t\17\2\2\u0109\u010a"+
		"\t\27\2\2\u010a\u010b\t\4\2\2\u010b@\3\2\2\2\u010c\u010d\5\t\5\2\u010d"+
		"\u010e\t\22\2\2\u010e\u010f\t\t\2\2\u010f\u0110\t\13\2\2\u0110\u0111\t"+
		"\f\2\2\u0111B\3\2\2\2\u0112\u0113\5\t\5\2\u0113\u0114\t\16\2\2\u0114\u0115"+
		"\t\7\2\2\u0115\u0116\t\17\2\2\u0116D\3\2\2\2\u0117\u0118\5\t\5\2\u0118"+
		"\u0119\t\7\2\2\u0119\u011a\t\n\2\2\u011a\u011b\t\f\2\2\u011b\u011c\t\23"+
		"\2\2\u011cF\3\2\2\2\u011d\u011e\5\t\5\2\u011e\u011f\t\7\2\2\u011fH\3\2"+
		"\2\2\u0120\u0121\5\t\5\2\u0121\u0122\t\17\2\2\u0122J\3\2\2\2\u0123\u0124"+
		"\t\b\2\2\u0124\u0125\t\b\2\2\u0125L\3\2\2\2\u0126\u0127\t\27\2\2\u0127"+
		"\u0128\t\b\2\2\u0128N\3\2\2\2\u0129\u012a\t\30\2\2\u012a\u012b\t\b\2\2"+
		"\u012bP\3\2\2\2\u012c\u012d\t\b\2\2\u012d\u012e\t\23\2\2\u012e\u012f\t"+
		"\6\2\2\u012f\u0130\t\f\2\2\u0130R\3\2\2\2\u0131\u0132\t\b\2\2\u0132\u0133"+
		"\t\n\2\2\u0133\u0134\t\r\2\2\u0134\u0135\t\r\2\2\u0135T\3\2\2\2\u0136"+
		"\u013a\5\5\3\2\u0137\u0139\n\31\2\2\u0138\u0137\3\2\2\2\u0139\u013c\3"+
		"\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2\2\2\u013b\u013d\3\2\2\2\u013c"+
		"\u013a\3\2\2\2\u013d\u013e\5\5\3\2\u013eV\3\2\2\2\u013f\u0141\t\32\2\2"+
		"\u0140\u013f\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0140\3\2\2\2\u0142\u0143"+
		"\3\2\2\2\u0143X\3\2\2\2\u0144\u0145\t\32\2\2\u0145\u0146\t\32\2\2\u0146"+
		"\u0148\t\32\2\2\u0147\u0149\t\33\2\2\u0148\u0147\3\2\2\2\u0149Z\3\2\2"+
		"\2\u014a\u014e\5\3\2\2\u014b\u014d\n\34\2\2\u014c\u014b\3\2\2\2\u014d"+
		"\u0150\3\2\2\2\u014e\u014c\3\2\2\2\u014e\u014f\3\2\2\2\u014f\u0152\3\2"+
		"\2\2\u0150\u014e\3\2\2\2\u0151\u014a\3\2\2\2\u0152\u0153\3\2\2\2\u0153"+
		"\u0151\3\2\2\2\u0153\u0154\3\2\2\2\u0154\\\3\2\2\2\u0155\u0157\t\35\2"+
		"\2\u0156\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158\u0156\3\2\2\2\u0158\u0159"+
		"\3\2\2\2\u0159^\3\2\2\2\u015a\u015e\5\7\4\2\u015b\u015d\n\36\2\2\u015c"+
		"\u015b\3\2\2\2\u015d\u0160\3\2\2\2\u015e\u015c\3\2\2\2\u015e\u015f\3\2"+
		"\2\2\u015f\u0164\3\2\2\2\u0160\u015e\3\2\2\2\u0161\u0163\t\36\2\2\u0162"+
		"\u0161\3\2\2\2\u0163\u0166\3\2\2\2\u0164\u0162\3\2\2\2\u0164\u0165\3\2"+
		"\2\2\u0165\u0167\3\2\2\2\u0166\u0164\3\2\2\2\u0167\u0168\b\60\2\2\u0168"+
		"`\3\2\2\2\u0169\u016b\t\37\2\2\u016a\u0169\3\2\2\2\u016b\u016c\3\2\2\2"+
		"\u016c\u016a\3\2\2\2\u016c\u016d\3\2\2\2\u016d\u016e\3\2\2\2\u016e\u016f"+
		"\b\61\2\2\u016fb\3\2\2\2\r\2\u013a\u0142\u0148\u014e\u0153\u0156\u0158"+
		"\u015e\u0164\u016c\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}