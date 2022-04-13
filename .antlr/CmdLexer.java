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
			"TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", "TOK_MKFS", "TOK_LOGIN", 
			"TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", "TOK_RMUSR", "TOK_MKFILE", 
			"TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", "TOK_PATH", "TOK_FIT", 
			"TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", "TOK_ID", "TOK_USUARIO", 
			"TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", "TOK_RUTA", "TOK_R", 
			"TOK_P", "TOK_FIRST", "TOK_WORST", "TOK_BEST", "TOK_KB", "TOK_MB", "TOK_BYTES", 
			"TOK_PRIMARIA", "TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FAST", "TOK_FULL", 
			"TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", "TOK_CAMINO", "TOK_PALABRA", 
			"TOK_IGUAL", "COMENTARIOS", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\64\u0173\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3\r\3"+
		"\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3"+
		"\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\22\3\22\3"+
		"\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3"+
		"\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26\3\26\3\27\3"+
		"\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3\32\3"+
		"\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3"+
		"\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3!\3!\3"+
		"!\3\"\3\"\3\"\3#\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3*"+
		"\3*\3*\3+\3+\3+\3+\3+\3,\3,\7,\u013f\n,\f,\16,\u0142\13,\3,\3,\3-\6-\u0147"+
		"\n-\r-\16-\u0148\3.\3.\3.\3.\3.\3/\3/\6/\u0152\n/\r/\16/\u0153\6/\u0156"+
		"\n/\r/\16/\u0157\3\60\6\60\u015b\n\60\r\60\16\60\u015c\3\61\3\61\3\62"+
		"\3\62\7\62\u0163\n\62\f\62\16\62\u0166\13\62\3\62\5\62\u0169\n\62\3\62"+
		"\3\62\3\63\6\63\u016e\n\63\r\63\16\63\u016f\3\63\3\63\2\2\64\3\3\5\4\7"+
		"\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64\3\2%\4\2OOoo\4\2MMmm\4"+
		"\2FFff\4\2KKkk\4\2UUuu\4\2TTtt\4\2HHhh\4\2QQqq\4\2WWww\4\2PPpp\4\2VVv"+
		"v\4\2NNnn\4\2IIii\4\2RRrr\4\2GGgg\4\2ZZzz\4\2EEee\4\2CCcc\4\2JJjj\4\2"+
		"\\\\||\4\2[[{{\4\2YYyy\4\2DDdd\3\2$$\5\2\f\f$$``\3\2\62;\5\2CC\\\\c|\3"+
		"\2\61\61\6\2\f\f\"\"$$``\7\2/\60\62;C\\aac|\3\2??\3\2%%\4\2\f\f``\3\2"+
		"\f\f\5\2\13\f\17\17\"\"\2\u017a\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2"+
		"\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2"+
		"\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2"+
		"\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2"+
		"\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2"+
		"\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2"+
		"\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O"+
		"\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2"+
		"\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3\2\2\2\3g\3\2\2\2"+
		"\5n\3\2\2\2\7u\3\2\2\2\t{\3\2\2\2\13\u0081\3\2\2\2\r\u0086\3\2\2\2\17"+
		"\u008c\3\2\2\2\21\u0093\3\2\2\2\23\u0099\3\2\2\2\25\u009f\3\2\2\2\27\u00a5"+
		"\3\2\2\2\31\u00ab\3\2\2\2\33\u00b2\3\2\2\2\35\u00b8\3\2\2\2\37\u00bd\3"+
		"\2\2\2!\u00c3\3\2\2\2#\u00c7\3\2\2\2%\u00cd\3\2\2\2\'\u00d2\3\2\2\2)\u00d8"+
		"\3\2\2\2+\u00de\3\2\2\2-\u00e4\3\2\2\2/\u00ea\3\2\2\2\61\u00ee\3\2\2\2"+
		"\63\u00f7\3\2\2\2\65\u0101\3\2\2\2\67\u0106\3\2\2\29\u010c\3\2\2\2;\u0111"+
		"\3\2\2\2=\u0117\3\2\2\2?\u011a\3\2\2\2A\u011d\3\2\2\2C\u0120\3\2\2\2E"+
		"\u0123\3\2\2\2G\u0126\3\2\2\2I\u0128\3\2\2\2K\u012a\3\2\2\2M\u012c\3\2"+
		"\2\2O\u012e\3\2\2\2Q\u0130\3\2\2\2S\u0132\3\2\2\2U\u0137\3\2\2\2W\u013c"+
		"\3\2\2\2Y\u0146\3\2\2\2[\u014a\3\2\2\2]\u0155\3\2\2\2_\u015a\3\2\2\2a"+
		"\u015e\3\2\2\2c\u0160\3\2\2\2e\u016d\3\2\2\2gh\t\2\2\2hi\t\3\2\2ij\t\4"+
		"\2\2jk\t\5\2\2kl\t\6\2\2lm\t\3\2\2m\4\3\2\2\2no\t\7\2\2op\t\2\2\2pq\t"+
		"\4\2\2qr\t\5\2\2rs\t\6\2\2st\t\3\2\2t\6\3\2\2\2uv\t\b\2\2vw\t\4\2\2wx"+
		"\t\5\2\2xy\t\6\2\2yz\t\3\2\2z\b\3\2\2\2{|\t\2\2\2|}\t\t\2\2}~\t\n\2\2"+
		"~\177\t\13\2\2\177\u0080\t\f\2\2\u0080\n\3\2\2\2\u0081\u0082\t\2\2\2\u0082"+
		"\u0083\t\3\2\2\u0083\u0084\t\b\2\2\u0084\u0085\t\6\2\2\u0085\f\3\2\2\2"+
		"\u0086\u0087\t\r\2\2\u0087\u0088\t\t\2\2\u0088\u0089\t\16\2\2\u0089\u008a"+
		"\t\5\2\2\u008a\u008b\t\13\2\2\u008b\16\3\2\2\2\u008c\u008d\t\r\2\2\u008d"+
		"\u008e\t\t\2\2\u008e\u008f\t\16\2\2\u008f\u0090\t\t\2\2\u0090\u0091\t"+
		"\n\2\2\u0091\u0092\t\f\2\2\u0092\20\3\2\2\2\u0093\u0094\t\2\2\2\u0094"+
		"\u0095\t\3\2\2\u0095\u0096\t\16\2\2\u0096\u0097\t\7\2\2\u0097\u0098\t"+
		"\17\2\2\u0098\22\3\2\2\2\u0099\u009a\t\7\2\2\u009a\u009b\t\2\2\2\u009b"+
		"\u009c\t\16\2\2\u009c\u009d\t\7\2\2\u009d\u009e\t\17\2\2\u009e\24\3\2"+
		"\2\2\u009f\u00a0\t\2\2\2\u00a0\u00a1\t\3\2\2\u00a1\u00a2\t\n\2\2\u00a2"+
		"\u00a3\t\6\2\2\u00a3\u00a4\t\7\2\2\u00a4\26\3\2\2\2\u00a5\u00a6\t\7\2"+
		"\2\u00a6\u00a7\t\2\2\2\u00a7\u00a8\t\n\2\2\u00a8\u00a9\t\6\2\2\u00a9\u00aa"+
		"\t\7\2\2\u00aa\30\3\2\2\2\u00ab\u00ac\t\2\2\2\u00ac\u00ad\t\3\2\2\u00ad"+
		"\u00ae\t\b\2\2\u00ae\u00af\t\5\2\2\u00af\u00b0\t\r\2\2\u00b0\u00b1\t\20"+
		"\2\2\u00b1\32\3\2\2\2\u00b2\u00b3\t\2\2\2\u00b3\u00b4\t\3\2\2\u00b4\u00b5"+
		"\t\4\2\2\u00b5\u00b6\t\5\2\2\u00b6\u00b7\t\7\2\2\u00b7\34\3\2\2\2\u00b8"+
		"\u00b9\t\20\2\2\u00b9\u00ba\t\21\2\2\u00ba\u00bb\t\20\2\2\u00bb\u00bc"+
		"\t\22\2\2\u00bc\36\3\2\2\2\u00bd\u00be\t\17\2\2\u00be\u00bf\t\23\2\2\u00bf"+
		"\u00c0\t\n\2\2\u00c0\u00c1\t\6\2\2\u00c1\u00c2\t\20\2\2\u00c2 \3\2\2\2"+
		"\u00c3\u00c4\t\7\2\2\u00c4\u00c5\t\20\2\2\u00c5\u00c6\t\17\2\2\u00c6\""+
		"\3\2\2\2\u00c7\u00c8\7/\2\2\u00c8\u00c9\t\17\2\2\u00c9\u00ca\t\23\2\2"+
		"\u00ca\u00cb\t\f\2\2\u00cb\u00cc\t\24\2\2\u00cc$\3\2\2\2\u00cd\u00ce\7"+
		"/\2\2\u00ce\u00cf\t\b\2\2\u00cf\u00d0\t\5\2\2\u00d0\u00d1\t\f\2\2\u00d1"+
		"&\3\2\2\2\u00d2\u00d3\7/\2\2\u00d3\u00d4\t\6\2\2\u00d4\u00d5\t\5\2\2\u00d5"+
		"\u00d6\t\25\2\2\u00d6\u00d7\t\20\2\2\u00d7(\3\2\2\2\u00d8\u00d9\7/\2\2"+
		"\u00d9\u00da\t\n\2\2\u00da\u00db\t\13\2\2\u00db\u00dc\t\5\2\2\u00dc\u00dd"+
		"\t\f\2\2\u00dd*\3\2\2\2\u00de\u00df\7/\2\2\u00df\u00e0\t\13\2\2\u00e0"+
		"\u00e1\t\23\2\2\u00e1\u00e2\t\2\2\2\u00e2\u00e3\t\20\2\2\u00e3,\3\2\2"+
		"\2\u00e4\u00e5\7/\2\2\u00e5\u00e6\t\f\2\2\u00e6\u00e7\t\26\2\2\u00e7\u00e8"+
		"\t\17\2\2\u00e8\u00e9\t\20\2\2\u00e9.\3\2\2\2\u00ea\u00eb\7/\2\2\u00eb"+
		"\u00ec\t\5\2\2\u00ec\u00ed\t\4\2\2\u00ed\60\3\2\2\2\u00ee\u00ef\7/\2\2"+
		"\u00ef\u00f0\t\n\2\2\u00f0\u00f1\t\6\2\2\u00f1\u00f2\t\n\2\2\u00f2\u00f3"+
		"\t\23\2\2\u00f3\u00f4\t\7\2\2\u00f4\u00f5\t\5\2\2\u00f5\u00f6\t\t\2\2"+
		"\u00f6\62\3\2\2\2\u00f7\u00f8\7/\2\2\u00f8\u00f9\t\17\2\2\u00f9\u00fa"+
		"\t\23\2\2\u00fa\u00fb\t\6\2\2\u00fb\u00fc\t\6\2\2\u00fc\u00fd\t\27\2\2"+
		"\u00fd\u00fe\t\t\2\2\u00fe\u00ff\t\7\2\2\u00ff\u0100\t\4\2\2\u0100\64"+
		"\3\2\2\2\u0101\u0102\7/\2\2\u0102\u0103\t\17\2\2\u0103\u0104\t\27\2\2"+
		"\u0104\u0105\t\4\2\2\u0105\66\3\2\2\2\u0106\u0107\7/\2\2\u0107\u0108\t"+
		"\22\2\2\u0108\u0109\t\t\2\2\u0109\u010a\t\13\2\2\u010a\u010b\t\f\2\2\u010b"+
		"8\3\2\2\2\u010c\u010d\7/\2\2\u010d\u010e\t\16\2\2\u010e\u010f\t\7\2\2"+
		"\u010f\u0110\t\17\2\2\u0110:\3\2\2\2\u0111\u0112\7/\2\2\u0112\u0113\t"+
		"\7\2\2\u0113\u0114\t\n\2\2\u0114\u0115\t\f\2\2\u0115\u0116\t\23\2\2\u0116"+
		"<\3\2\2\2\u0117\u0118\7/\2\2\u0118\u0119\t\7\2\2\u0119>\3\2\2\2\u011a"+
		"\u011b\7/\2\2\u011b\u011c\t\17\2\2\u011c@\3\2\2\2\u011d\u011e\t\b\2\2"+
		"\u011e\u011f\t\b\2\2\u011fB\3\2\2\2\u0120\u0121\t\27\2\2\u0121\u0122\t"+
		"\b\2\2\u0122D\3\2\2\2\u0123\u0124\t\30\2\2\u0124\u0125\t\b\2\2\u0125F"+
		"\3\2\2\2\u0126\u0127\t\3\2\2\u0127H\3\2\2\2\u0128\u0129\t\2\2\2\u0129"+
		"J\3\2\2\2\u012a\u012b\t\30\2\2\u012bL\3\2\2\2\u012c\u012d\t\17\2\2\u012d"+
		"N\3\2\2\2\u012e\u012f\t\20\2\2\u012fP\3\2\2\2\u0130\u0131\t\r\2\2\u0131"+
		"R\3\2\2\2\u0132\u0133\t\b\2\2\u0133\u0134\t\23\2\2\u0134\u0135\t\6\2\2"+
		"\u0135\u0136\t\f\2\2\u0136T\3\2\2\2\u0137\u0138\t\b\2\2\u0138\u0139\t"+
		"\n\2\2\u0139\u013a\t\r\2\2\u013a\u013b\t\r\2\2\u013bV\3\2\2\2\u013c\u0140"+
		"\t\31\2\2\u013d\u013f\t\32\2\2\u013e\u013d\3\2\2\2\u013f\u0142\3\2\2\2"+
		"\u0140\u013e\3\2\2\2\u0140\u0141\3\2\2\2\u0141\u0143\3\2\2\2\u0142\u0140"+
		"\3\2\2\2\u0143\u0144\t\31\2\2\u0144X\3\2\2\2\u0145\u0147\t\33\2\2\u0146"+
		"\u0145\3\2\2\2\u0147\u0148\3\2\2\2\u0148\u0146\3\2\2\2\u0148\u0149\3\2"+
		"\2\2\u0149Z\3\2\2\2\u014a\u014b\t\33\2\2\u014b\u014c\t\33\2\2\u014c\u014d"+
		"\t\33\2\2\u014d\u014e\t\34\2\2\u014e\\\3\2\2\2\u014f\u0151\t\35\2\2\u0150"+
		"\u0152\t\36\2\2\u0151\u0150\3\2\2\2\u0152\u0153\3\2\2\2\u0153\u0151\3"+
		"\2\2\2\u0153\u0154\3\2\2\2\u0154\u0156\3\2\2\2\u0155\u014f\3\2\2\2\u0156"+
		"\u0157\3\2\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158^\3\2\2\2"+
		"\u0159\u015b\t\37\2\2\u015a\u0159\3\2\2\2\u015b\u015c\3\2\2\2\u015c\u015a"+
		"\3\2\2\2\u015c\u015d\3\2\2\2\u015d`\3\2\2\2\u015e\u015f\t \2\2\u015fb"+
		"\3\2\2\2\u0160\u0164\t!\2\2\u0161\u0163\t\"\2\2\u0162\u0161\3\2\2\2\u0163"+
		"\u0166\3\2\2\2\u0164\u0162\3\2\2\2\u0164\u0165\3\2\2\2\u0165\u0168\3\2"+
		"\2\2\u0166\u0164\3\2\2\2\u0167\u0169\t#\2\2\u0168\u0167\3\2\2\2\u0168"+
		"\u0169\3\2\2\2\u0169\u016a\3\2\2\2\u016a\u016b\b\62\2\2\u016bd\3\2\2\2"+
		"\u016c\u016e\t$\2\2\u016d\u016c\3\2\2\2\u016e\u016f\3\2\2\2\u016f\u016d"+
		"\3\2\2\2\u016f\u0170\3\2\2\2\u0170\u0171\3\2\2\2\u0171\u0172\b\63\2\2"+
		"\u0172f\3\2\2\2\13\2\u0140\u0148\u0153\u0157\u015c\u0164\u0168\u016f\3"+
		"\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}